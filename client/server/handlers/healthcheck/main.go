package healthcheck

import (
	"grpc-service/client/config"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	log *logrus.Entry
}

func New(cfg config.Config) *Handler {
	return &Handler{
		log: cfg.Log(),
	}
}
