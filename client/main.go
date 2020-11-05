package app

import (
	"fmt"
	"grpc-service/client/config"
	"grpc-service/client/server"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type App struct {
	config config.Config
	log    *logrus.Entry
}

func New(config config.Config) *App {
	return &App{
		config: config,
		log:    config.Log(),
	}
}

func (a *App) Start() error {
	conf := a.config

	httpConfiguration := conf.HTTP()
	router := server.Router(conf)

	serverHost := fmt.Sprintf("%s:%s", httpConfiguration.Host, httpConfiguration.Port)
	a.log.WithField("api", "start").
		Info(fmt.Sprintf("listenig addr =  %s", serverHost))

	httpServer := http.Server{
		Addr:           serverHost,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}

	return nil
}
