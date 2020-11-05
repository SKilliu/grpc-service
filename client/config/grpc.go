package config

import "github.com/caarlos0/env"

type GRPCClient struct {
	Host string `env:"GRPC_CLIENT_HOST,required"`
	Port string `env:"GRPC_CLIENT_PORT,required"`
}

func (c *ConfigImpl) GRPCClient() *GRPCClient {
	if c.grpc != nil {
		return c.grpc
	}

	c.Lock()
	defer c.Unlock()

	grpc := &GRPCClient{}
	if err := env.Parse(grpc); err != nil {
		panic(err)
	}

	c.grpc = grpc

	return c.grpc
}
