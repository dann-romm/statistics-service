package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"statistics-service/internal/service"
)

func NewRouter(handler *echo.Echo, services *service.Service) {
	handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}", "method":"${method}","uri":"${uri}", "status":${status},"error":"${error}"}` + "\n",
		Output: setLogsFile(),
	}))

	handler.GET("/healthz", func(c echo.Context) error { return c.String(200, "OK") })

	auth := handler.Group("/auth")
	{
		newAuthRoutes(auth, services)
	}

	authMiddleware := &AuthMiddleware{services.Auth}
	v1 := handler.Group("/api/v1", authMiddleware.UserIdentity)
	{
		newStatisticsRoutes(v1, services)
	}
}

func setLogsFile() *os.File {
	file, err := os.OpenFile("logs/v1.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
