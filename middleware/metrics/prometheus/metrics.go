package prometheus

import (
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var defaultLabelNames = []string{"node", "host", "uri", "method", "status"}

func MetricsFunc(options ...Option) echo.MiddlewareFunc {
	opts := applyOptions(options...)

	reqCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Name:      "request_total",
			Help:      "Total request count.",
		},
		defaultLabelNames,
	)

	reqDur := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Name:      "request_duration",
			Help:      "Request duration in nanoseconds.",
		},
		defaultLabelNames,
	)

	reqSize := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Name:      "request_size",
			Help:      "Request size in bytes.",
		},
		defaultLabelNames,
	)

	resSize := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Name:      "response_size",
			Help:      "Response size in bytes.",
		},
		defaultLabelNames,
	)

	opts.Registry.MustRegister(reqCount, reqDur, reqSize, resSize)

	hostname, err := os.Hostname()
	if err != nil {
		log.Warnf("os.Hostname() error:%v", err)
		hostname = "-"
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			// 拦截metrics path，默认"/metrics"
			if req.URL.Path == opts.MetricsPath {
				promhttp.Handler().ServeHTTP(c.Response(), c.Request())
				return nil
			}

			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}

			latency := time.Since(start)
			status := strconv.Itoa(res.Status)

			reqCount.WithLabelValues(hostname, req.Host, req.RequestURI, req.Method, status).Inc()
			reqDur.WithLabelValues(hostname, req.Host, req.RequestURI, req.Method, status).Observe(float64(latency.Nanoseconds()))
			reqSize.WithLabelValues(hostname, req.Host, req.RequestURI, req.Method, status).Observe(float64(req.ContentLength))
			resSize.WithLabelValues(hostname, req.Host, req.RequestURI, req.Method, status).Observe(float64(res.Size))

			return nil
		}
	}
}
