package server

import (
	"fmt"
	"grpc-service/client/config"
	"grpc-service/client/server/handlers/coordinates"
	"log"

	"github.com/pkg/errors"

	"google.golang.org/grpc"

	"github.com/SKilliu/grpc-service/proto/protogo"
)

// Provider persist handlers service structures.
type Provider struct {
	coordinates *coordinates.Handler
}

// NewProvider is provider constructor.
func NewProvider(cfg config.Config) *Provider {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.GRPCClient().Host, cfg.GRPCClient().Port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create connection for grpc client"))
	}

	client := protogo.NewCoordinatesSaverClient(conn)
	return &Provider{
		coordinates: coordinates.New(cfg.Log(), client),
	}
}
