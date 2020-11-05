package config

import (
	"io/ioutil"
	"os"
	"sync"

	"gopkg.in/yaml.v2"

	"github.com/sirupsen/logrus"
)

type Config interface {
	HTTP() *HTTP
	Log() *logrus.Entry
	GRPCClient() *GRPCClient
}

type ConfigImpl struct {
	sync.Mutex

	http *HTTP
	log  *logrus.Entry
	grpc *GRPCClient
}

func New() Config {
	return &ConfigImpl{
		Mutex: sync.Mutex{},
	}
}

func UploadEnvironmentVariables(pathToConfigFile string) {
	ymlFile, err := os.Open(pathToConfigFile)
	if err != nil {
		panic(err)
	}

	defer ymlFile.Close()

	var variables = make(map[string]string)

	byteValue, err := ioutil.ReadAll(ymlFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(byteValue, &variables)
	if err != nil {
		panic(err)
	}

	for k, v := range variables {
		err := os.Setenv(k, v)
		if err != nil {
			panic(err)
		}
	}
}
