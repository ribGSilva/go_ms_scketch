package config

import (
	"fmt"
	"gabrielrsilva/go-ms-scketch/constant"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ConfigLogger(router *gin.Engine) {
	//log.SetFormatter(&log.JSONFormatter{})

	router.Use(newRequestLogger())
}

func newRequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info(fmt.Sprintf("%s to endpoint %s with traceid %s",
			ctx.Request.Method,
			ctx.Request.RequestURI,
			ctx.Request.Header.Get(constant.REQUEST_TRACE_ID)),
		)
	}
}

func LogResponse(trace string, method string, path string, status int, response interface{}) {
	log.Info(fmt.Sprintf("%s to endpoint %s with traceid %s ended with status %d and response %+v",
		method,
		path,
		trace,
		status,
		response,
	))
}
