package router

import (
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"

	"github.com/yuleihua/trade/conf"
	"github.com/yuleihua/trade/middleware/opentracing"
	"github.com/yuleihua/trade/transfer/api"
)

func Router(c *conf.Config) *echo.Echo {
	// Echo instance
	e := echo.New()

	// Customization
	if c.Cmn.IsRelease {
		e.Debug = false
	}
	e.Logger.SetPrefix("v1")
	e.Logger.SetLevel(2)

	// OpenTracing
	if c.OpenTracing.IsOnline {
		e.Use(opentracing.OpenTracing("api"))
	}

	// Gzip
	e.Use(mw.GzipWithConfig(mw.GzipConfig{Level: 5}))

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routers
	e.POST("/login", api.CustomerLoginHandler)
	e.GET("/ping", api.PingHandler)

	// JWT
	r := e.Group("/v1")
	r.Use(mw.JWTWithConfig(mw.JWTConfig{
		SigningKey:  []byte("hello-88773dy2"),
		ContextKey:  "_user",
		TokenLookup: "header:" + echo.HeaderAuthorization,
	}))

	// curl http://localhost/v1/transfers -H "Authorization: Bearer XXX"
	r.GET("/transfers", api.TransferDetailHandler)
	r.POST("/add-transfer", api.AddTransferHandler)
	r.POST("/confirm-transfer", api.ConfirmTransferHandler)

	return e
}
