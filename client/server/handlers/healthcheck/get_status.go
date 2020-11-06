package healthcheck

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusResp struct {
	Status string `json:"status"`
}

func (h *Handler) GetStatus(c echo.Context) error {
	h.log.Info("client API is working")

	body, _ := json.Marshal(StatusResp{Status: "ok"})

	return c.JSON(http.StatusOK, body)
}
