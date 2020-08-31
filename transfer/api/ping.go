package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func PingHandler(c echo.Context) error {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": "pong",
	})
	return nil
}
