package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusResp struct {
	Status string `json:"status"`
} //@name StatusResp

// GetStatus godoc
// @Summary Healthcheck
// @Tags healthcheck
// @Consume application/json
// @Description Check connection to client
// @Accept  json
// @Produce json
// @Success 200 {object} StatusResp
// @Router /status [get]
func (h *Handler) GetStatus(c echo.Context) error {
	h.log.Info("client API is working")

	return c.JSON(http.StatusOK, StatusResp{
		Status: "ok",
	})
}
