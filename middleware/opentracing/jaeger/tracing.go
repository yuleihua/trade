package jaeger

import (
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
)

type jaegerLogger struct{}

func (l *jaegerLogger) Error(msg string) {
	log.Debugf("ERROR: %s", msg)
}

func (l *jaegerLogger) Infof(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}

func SetupTracer(serviceName, addr string) io.Closer {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := &jaegerLogger{}
	jMetricsFactory := metrics.NullFactory

	metricsFactory := metrics.NewLocalFactory(0)
	metrics := jaeger.NewMetrics(metricsFactory, nil)

	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		log.Printf("could not initialize jaeger sender: %s", err.Error())
		return nil
	}

	repoter := jaeger.NewRemoteReporter(sender, jaeger.ReporterOptions.Metrics(metrics))

	// Initialize tracer with a logger and a metrics factory
	closer, err := cfg.InitGlobalTracer(
		serviceName,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
		jaegercfg.Reporter(repoter),
	)

	if err != nil {
		log.Printf("could not initialize jaeger tracer: %s", err.Error())
		return nil
	}
	return closer
}
