package main

import (
	"gabrielrsilva/go-ms-scketch/config"
	"gabrielrsilva/go-ms-scketch/endpoint"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(config.NewRestLogger())
	endpoint.InjectPingEndpoints(router)
	endpoint.InjectHelloEndpoit(router)

	router.Run(":8080")
}
