package coordinates

import (
	"context"
	"grpc-service/client/server/errs"
	"net/http"

	"github.com/SKilliu/grpc-service/proto/protogo"
	"github.com/labstack/echo/v4"
)

type SaveCoordinatesReq struct {
	Location  string  `json:"location"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func (h *Handler) SaveCoordinates(c echo.Context) error {
	var req SaveCoordinatesReq

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to decode add device id request")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	resp, err := h.grpcClient.SaveCoordinates(context.Background(), &protogo.SaveRequest{
		Location:  req.Location,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	})
	if err != nil {
		h.log.WithError(err).Error("failed to send request to grpc server")
		return c.JSON(http.StatusBadRequest, errs.InternalServerErr)
	}

	return c.JSON(http.StatusOK, resp)
}
