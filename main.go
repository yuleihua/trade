package main

import (
	"context"
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/yuleihua/trade/conf"
	"github.com/yuleihua/trade/middleware/cors"
	"github.com/yuleihua/trade/middleware/metrics/prometheus"
	"github.com/yuleihua/trade/middleware/opentracing"
	"github.com/yuleihua/trade/middleware/pprof"
	"github.com/yuleihua/trade/pkg/logger"
	"github.com/yuleihua/trade/pkg/server"
	"github.com/yuleihua/trade/router"
	"github.com/yuleihua/trade/transfer/service"
)

const Version = "0.1.1"

var (
	confFile  string
	isVersion bool

	gitCommit string // commit hash
	buildDate string // build datetime
)

func init() {
	flag.StringVar(&confFile, "c", "config.yaml", "config file")
	flag.BoolVar(&isVersion, "v", false, "version information")
}

func versionInfo(name, commit, buildDate string) {
	fmt.Println("Version:", Version)
	if commit != "" {
		fmt.Println("Git Commit:", commit)
	}

	if buildDate != "" {
		fmt.Println("Build Date:", buildDate)
	}

	if name != "" {
		fmt.Println("Binary Name:", name)
	}

	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
}

func main() {
	flag.Parse()

	if isVersion {
		versionInfo(os.Args[0], gitCommit, buildDate)
		return
	}
	log.Debugf("running with conf:%s", confFile)

	// 配置初始化
	config, err := conf.Setup(confFile)
	if err != nil {
		log.Panic(err)
	}

	// setting the maximum number of CPUs.
	if config.Cmn.CPUs > 0 {
		runtime.GOMAXPROCS(config.Cmn.CPUs)
	}

	// setting logger
	logger.SetLogger(config.Cmn.LogPath, config.Cmn.LogFile, config.Cmn.LogLevel)

	// setting db
	t := service.NewTransfer()
	if err := t.Init(config); err != nil {
		log.Fatal(err)
	}

	// register transfer service
	server.Register(t.Name(), t)

	// echo init
	e := echo.New()

	// pprof
	if !config.Cmn.IsRelease {
		e.Pre(pprof.Serve())
	}

	e.Pre(mw.RemoveTrailingSlash())

	if config.OpenTracing.IsOnline {
		// OpenTracing
		traceCfg := opentracing.Configuration{
			IsOnline: config.OpenTracing.IsOnline,
			Type:     opentracing.TracerType(config.OpenTracing.Type),
		}
		if closer := traceCfg.InitGlobalTracer(
			opentracing.ServiceName(config.OpenTracing.ServiceName),
			opentracing.Address(config.OpenTracing.Address),
		); closer != nil {
			defer closer.Close()
		}
	}

	// 日志级别
	e.Logger.SetLevel(2)

	// Metrics
	if config.Metric.IsOnline {
		e.Use(prometheus.MetricsFunc(
			prometheus.Namespace(config.Server.AppName),
		))
	}

	//// Secure
	e.Use(mw.SecureWithConfig(mw.DefaultSecureConfig))
	e.Use(mw.MethodOverride())

	// CORS
	e.Use(mw.CORSWithConfig(cors.SetCors(config.Server.Origins)))

	r := router.Router(config)

	// Start server
	go func() {
		if err := r.Start(config.Server.Addr); err != nil {
			r.Logger.Errorf("Shutting down the server with error:%v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	chanExit := make(chan os.Signal, 1)
	signal.Notify(chanExit, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-chanExit:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := r.Shutdown(ctx); err != nil {
			r.Logger.Fatal(err)
		}
	}

	log.Warn("application stop")
}
