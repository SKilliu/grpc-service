package main

import (
	app "github.com/SKilliu/grpc-service"
	"github.com/SKilliu/grpc-service/config"

	"github.com/pkg/errors"
)

const pathToConfigFile = "./static/envs.yaml"

func main() {
	apiConfig := config.New()
	log := apiConfig.Log()

	//docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", apiConfig.HTTP().Host, apiConfig.HTTP().Port)
	config.UploadEnvironmentVariables(pathToConfigFile)
	api := app.New(apiConfig)
	if err := api.Start(); err != nil {
		log.WithError(err)
		panic(errors.Wrap(err, "failed to start grpc-server"))
	}
}
