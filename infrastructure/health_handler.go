package infrastructure

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthCheckHandler struct{

}

func NewHealthCheckHandler() *healthCheckHandler{
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) HealthCheck (c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}