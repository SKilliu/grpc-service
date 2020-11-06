package server

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	"grpc-service/client/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	router.GET("/status", provider.healthcheck.GetStatus)
	router.POST("/coordinates", provider.coordinates.SaveCoordinates)

	return router
}
