package main

import (
	_infra "clone-instagram-service/infrastructure"
	"clone-instagram-service/util"
	"fmt"
)


func main() {
	fmt.Println("Hello")

	
	e := util.InitEchoApp()

	healthCheckHandler := _infra.NewHealthCheckHandler()

	e.GET("/health", healthCheckHandler.HealthCheck)

	e.Logger.Fatal(e.Start(":5000"))
}


