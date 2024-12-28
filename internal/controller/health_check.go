package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthCheckController struct {
	server *echo.Echo
}

func NewHealthCheckController(server *echo.Echo) *HealthCheckController {
	return &HealthCheckController{server: server}
}

func (c *HealthCheckController) RegisterHandler() {
	group := c.server.Group("/health-check")
	group.GET("", c.Get)
}

func (c *HealthCheckController) Get(e echo.Context) error {
	return e.JSONPretty(http.StatusOK, nil, "")
}
