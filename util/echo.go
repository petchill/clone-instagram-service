package util

import "github.com/labstack/echo/v4"

func InitEchoApp() *echo.Echo{
	e := echo.New()
	return e
}