package coordinates

import (
	"context"
	"grpc-service/client/server/errs"
	"net/http"

	"github.com/SKilliu/grpc-service/proto/protogo"
	"github.com/labstack/echo/v4"
)

type SaveCoordinatesReq struct {
	Location  string  `json:"location" example:"London"`
	Latitude  float32 `json:"latitude" example:"12.345"`
	Longitude float32 `json:"longitude" example:"12.345"`
} //@name SaveCoordinatesReq

// Coordinates godoc
// @Summary Send coordinates to grpc server
// @Tags coordinates
// @Consume application/json
// @Param JSON body SaveCoordinatesReq true "Body for save coordinates request"
// @Description Send coordinates to grpc server for saving in a database
// @Accept  json
// @Produce  json
// @Success 200 {} http.StatusOK
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /coordinates [post]
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
