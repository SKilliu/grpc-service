package main

import (
	"github.com/SKilliu/grpc-service/server/config"

	app "github.com/SKilliu/grpc-service/server"

	"github.com/pkg/errors"
)

func main() {
	apiConfig := config.New()
	log := apiConfig.Log()

	//docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", apiConfig.HTTP().Host, apiConfig.HTTP().Port)

	api := app.New(apiConfig)
	if err := api.Start(); err != nil {
		log.WithError(err)
		panic(errors.Wrap(err, "failed to start grpc-server"))
	}
}
