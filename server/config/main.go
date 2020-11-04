package config

import (
	"grpc-service/server/db"
	"io/ioutil"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config interface {
	GRPCServer() *GRPCServer
	Log() *logrus.Entry
	DB() db.QInterface
}

type ConfigImpl struct {
	sync.Mutex

	grpcServer *GRPCServer
	log        *logrus.Entry
	db         db.QInterface
}

// New config creating
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
