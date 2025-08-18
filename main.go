package main

import (
	_infra "clone-instagram-service/internal/infrastructure"
	"clone-instagram-service/internal/util"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	e := util.InitEchoApp()

	healthCheckHandler := _infra.NewHealthCheckHandler()

	e.GET("/health", healthCheckHandler.HealthCheck)

	e.Logger.Fatal(e.Start(":5000"))
}
