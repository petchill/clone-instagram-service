package main

import (
	_infra "clone-instagram-service/internal/infrastructure"
	_handler "clone-instagram-service/internal/infrastructure/handler"
	"clone-instagram-service/internal/util"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	e := util.InitEchoApp()

	mediaHandler := _handler.NewMediaHandler()
	healthCheckHandler := _infra.NewHealthCheckHandler()

	mediaHandler.RegisterRoutes(e)

	e.GET("/health", healthCheckHandler.HealthCheck)

	e.Logger.Fatal(e.Start(":5000"))
}
