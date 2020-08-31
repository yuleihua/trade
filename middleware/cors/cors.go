package cors

import (
	"net/http"

	"github.com/labstack/echo/v4/middleware"
)

func SetCors(origins []string) middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins: origins,
		AllowHeaders: []string{"*"},
		AllowMethods: []string{
			http.MethodOptions,
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	}
}
