package config

import (
	"context"

	"github.com/SKilliu/grpc-service/proto/protogo"

	"github.com/caarlos0/env"
)

type GRPCServer struct {
	Host string `env:"GRPC_SERVER_HOST,required"`
	Port string `env:"GRPC_SERVER_PORT,required"`
}

func (s *GRPCServer) SaveCoordinates(ctx context.Context, req *protogo.SaveRequest) (*protogo.SaveResponse, error) {

	return &protogo.SaveResponse{OperationResult: "ok"}, nil
}

func (c *ConfigImpl) GRPCServer() *GRPCServer {
	if c.grpcServer != nil {
		return c.grpcServer
	}

	c.Lock()
	defer c.Unlock()

	grpcServer := &GRPCServer{}
	if err := env.Parse(grpcServer); err != nil {
		panic(err)
	}

	c.grpcServer = grpcServer

	return c.grpcServer
}
