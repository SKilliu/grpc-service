package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusResp struct {
	Status string `json:"status"`
}

func (h *Handler) GetStatus(c echo.Context) error {
	h.log.Info("client API is working")

	return c.JSON(http.StatusOK, StatusResp{
		Status: "ok",
	})
}
