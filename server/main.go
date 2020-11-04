package app

import (
	"fmt"
	"net"

	"github.com/SKilliu/grpc-service/proto/protogo"

	"github.com/SKilliu/grpc-service/server/config"

	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
)

// App is store necessary config and logger entity for app creation and starting
type App struct {
	config config.Config
	log    *logrus.Entry
}

// New app config creating
func New(config config.Config) *App {
	return &App{
		config: config,
		log:    config.Log(),
	}
}

// Start the app
func (a *App) Start() error {
	conf := a.config

	serverConfiguration := conf.GRPCServer()

	serverHost := fmt.Sprintf("%s:%s", serverConfiguration.Host, serverConfiguration.Port)
	a.log.WithField("grpc-server", "start").
		Info(fmt.Sprintf("listenig addr =  %s", serverHost))

	s := grpc.NewServer()
	srv := &config.GRPCServer{}

	protogo.RegisterCoordinatesSaverServer(s, srv)
	l, err := net.Listen("tcp", serverHost)
	if err != nil {
		return err
	}

	if err := s.Serve(l); err != nil {
		return err
	}

	return nil
}
