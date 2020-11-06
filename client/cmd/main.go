package main

import (
	"fmt"
	app "grpc-service/client"
	"grpc-service/client/config"
	"grpc-service/client/docs"

	"github.com/pkg/errors"
)

const pathToConfigFile = "./static/envs.yaml"

// @title grpc-service
// @version 1.0
func main() {
	apiConfig := config.New()
	log := apiConfig.Log()

	config.UploadEnvironmentVariables(pathToConfigFile)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", apiConfig.HTTP().Host, apiConfig.HTTP().Port)
	api := app.New(apiConfig)
	if err := api.Start(); err != nil {
		log.WithError(err)
		panic(errors.Wrap(err, "failed to start api server"))
	}
}
