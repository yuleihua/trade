package logger

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

func newRotateHook(logFileName string, maxAge time.Duration, rotationTime time.Duration) *lfshook.LfsHook {
	writer, err := rotatelogs.New(
		logFileName+".%Y%m%d",
		rotatelogs.WithLinkName(logFileName),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		log.Fatalf("config logger error:%v", err)
	}
	return lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true, TimestampFormat: "2006/01/02-15:04:05.000"})
}

func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Errorf("set null err:%v", err)
	}
	writer := bufio.NewWriter(src)
	log.SetOutput(writer)
}

func SetLogger(logPath, logFile string, logLevel int) {

	// setting logger
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006/01/02-15:04:05.000",
	})

	var logfile string
	if logFile == "console" || logFile == "" {
		logfile = "console"
	} else if logFile != "" {
		if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
			fmt.Printf("create log directory error:%v, path:%\n", err, logPath)
			os.Exit(1)
		}
		logfile = filepath.Join(logPath, logFile)
	}

	level := logLevel
	if logLevel <= 0 {
		level = 5
	}
	log.SetLevel(log.Level(level))
	log.SetReportCaller(true)

	if logfile != "" && logfile != "console" {
		if err := os.Remove(logfile); err != nil {
			log.Errorf("remove log file, error:%v", err)
		}
		log.AddHook(newRotateHook(logfile, 3*24*time.Hour, 24*time.Hour))
		//log.AddHook(newRotateHook(logfile, 3*time.Hour, time.Hour))
		setNull()
	}
}
