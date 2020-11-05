package server

import (
	"context"

	"github.com/SKilliu/grpc-service/config"

	"github.com/google/uuid"

	"github.com/SKilliu/grpc-service/db/models"

	"github.com/SKilliu/grpc-service/db"
	"github.com/sirupsen/logrus"

	"github.com/SKilliu/grpc-service/proto/protogo"
)

type GRPCServer struct {
	log           *logrus.Entry
	coordinatesDB db.CoordinatesQ
}

func New(cfg config.Config) *GRPCServer {
	return &GRPCServer{
		log:           cfg.Log(),
		coordinatesDB: cfg.DB().CoordinatesQ(),
	}
}

func (s *GRPCServer) SaveCoordinates(ctx context.Context, req *protogo.SaveRequest) (*protogo.SaveResponse, error) {
	err := s.coordinatesDB.Insert(models.Coordinates{
		ID:        uuid.New().String(),
		Location:  req.Location,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	})
	if err != nil {
		s.log.WithError(err).Error("failed to insert coordinates to db")
		return &protogo.SaveResponse{
			OperationResult: "Operation failed!",
		}, err
	}

	s.log.Info(req)

	return &protogo.SaveResponse{OperationResult: "Operation success"}, nil
}
