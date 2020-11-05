package coordinates

import (
	"github.com/SKilliu/grpc-service/proto/protogo"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	log        *logrus.Entry
	grpcClient protogo.CoordinatesSaverClient
}

func New(log *logrus.Entry, client protogo.CoordinatesSaverClient) *Handler {
	return &Handler{
		log:        log,
		grpcClient: client,
	}
}
