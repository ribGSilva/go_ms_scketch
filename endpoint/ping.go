package endpoint

import (
	"gabrielrsilva/go-ms-scketch/config"
	"gabrielrsilva/go-ms-scketch/constant"
	"gabrielrsilva/go-ms-scketch/service"

	"github.com/gin-gonic/gin"
)

func InjectPingEndpoints(router *gin.Engine) {
	router.GET("/ping", getPing)
}

func getPing(ctx *gin.Context) {
	service := service.GetPingService()
	response := service.GetPing()
	config.LogResponse(
		ctx.GetHeader(constant.REQUEST_TRACE_ID),
		ctx.Request.Method,
		ctx.Request.URL.Path,
		response.GetStatus(),
		response.GetResponse(),
	)
	if response.HasResponse() {
		ctx.JSON(response.GetStatus(), response.GetResponse())
	} else {
		ctx.Status(response.GetStatus())
	}
}
