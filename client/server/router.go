package server

import (
	"grpc-service/client/config"
	"grpc-service/client/server/middlewares"

	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

const durationThreshold = time.Second * 10

func Router(cfg config.Config) *echo.Echo {
	router := echo.New()

	cors := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*", "GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"*", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-auth", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"},
		ExposeHeaders:    []string{"*", "Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	provider := NewProvider(cfg)

	router.Use(
		cors,
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.DefaultLoggerConfig),
	)

	router.GET("/status", provider.healthcheck.Status)
	router.POST("/connect", provider.coordinates.SaveCoordinates)

	return router
}
