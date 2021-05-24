package main

import (
	"gabrielrsilva/go-ms-scketch/config"
	"gabrielrsilva/go-ms-scketch/endpoint"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	config.ConfigRouter(router)
	config.ConfigLogger(router)
	endpoint.InjectPingEndpoints(router)

	router.Run()
}
