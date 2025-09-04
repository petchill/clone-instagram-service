package util

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitEchoApp() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	return e
}
