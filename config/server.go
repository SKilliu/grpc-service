package config

import (
	"github.com/caarlos0/env"
)

type GRPCServer struct {
	Host string `env:"GRPC_SERVER_HOST,required"`
	Port string `env:"GRPC_SERVER_PORT,required"`
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
