package coordinates

import (
	"context"
	"encoding/json"
	"grpc-service/client/server/errs"
	"net/http"

	"github.com/SKilliu/grpc-service/proto/protogo"
)

type SaveCoordinatesReq struct {
	Location  string  `json:"location"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func (h *Handler) SaveCoordinates(w http.ResponseWriter, r *http.Request) {
	var req SaveCoordinatesReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to decode add device id request")
		errs.BadRequest(w, errs.BadParamInBodyErr)
		return
	}

	resp, err := h.grpcClient.SaveCoordinates(context.Background(), &protogo.SaveRequest{
		Location:  req.Location,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	})
	if err != nil {
		h.log.WithError(err).Error("failed to send request to grpc server")
		errs.InternalError(w)
		return
	}

	serializedBody, err := json.Marshal(resp)
	if err != nil {
		h.log.WithError(err).Error("failed to marshal get item request")
		errs.InternalError(w)
		return
	}

	_, _ = w.Write(serializedBody)
}
