package server

import (
	"grpc-service/client/config"
	"grpc-service/client/server/middlewares"

	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

const durationThreshold = time.Second * 10

func Router(cfg config.Config) chi.Router {
	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*", "GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-auth", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"},
		ExposedHeaders:   []string{"*", "Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	provider := NewProvider(cfg)

	router.Use(
		cors.Handler,
		middleware.Recoverer,
		middlewares.Logger(cfg.Log(), durationThreshold),
	)

	router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello gRPC Client API"))
	})

	router.Post("/coordinates", provider.coordinates.SaveCoordinates)

	return router
}
