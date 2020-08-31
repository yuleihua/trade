package opentracing

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"

	"github.com/yuleihua/trade/middleware/opentracing/jaeger"
)

type TracerType string

const (
	DefaultKey              = "trade_key"
	TracerJaeger TracerType = "jaeger"
)

type Configuration struct {
	IsOnline bool
	Type     TracerType
}

func (c Configuration) InitGlobalTracer(options ...Option) io.Closer {
	if c.IsOnline {
		return nil
	} else {
		opts := applyOptions(c.Type, options...)

		switch c.Type {
		case TracerJaeger:
			return jaeger.SetupTracer(opts.ServiceName, opts.Address)
		default:
			return nil
		}
	}
}

func OpenTracing(comp string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var span opentracing.Span
			opName := comp + ":" + c.Request().URL.Path
			// 检查Header中是否有Trace信息
			wireContext, err := opentracing.GlobalTracer().Extract(
				opentracing.TextMap,
				opentracing.HTTPHeadersCarrier(c.Request().Header))
			if err != nil {
				// 启动新Span
				span = opentracing.StartSpan(opName)
			} else {
				log.Debugf("opentracing span child!")
				span = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
			}

			defer span.Finish()
			c.Set(DefaultKey, span)

			span.SetTag("component", comp)
			span.SetTag("span.kind", "server")
			span.SetTag("http.url", c.Request().Host+c.Request().RequestURI)
			span.SetTag("http.method", c.Request().Method)

			if err := next(c); err != nil {
				span.SetTag("error", true)
				c.Error(err)
			}

			span.SetTag("error", false)
			span.SetTag("http.status_code", c.Response().Status)

			return nil
		}
	}
}

func Default(c echo.Context) opentracing.Span {
	ot := c.Get(DefaultKey)
	if ot == nil {
		return nil
	}
	return c.Get(DefaultKey).(opentracing.Span)
}
