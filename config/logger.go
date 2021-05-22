package config

import (
	"gabrielrsilva/go-ms-scketch/constant"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ConfigLogger(router *gin.Engine) {
	log.SetFormatter(&log.JSONFormatter{})

	router.Use(newRequestLogger())
}

func newRequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info("%s: %s to endpoint %s with traceid %s",
			time.Now().Format(time.RFC3339),
			ctx.Request.Method,
			ctx.Request.RequestURI,
			ctx.Request.Header.Get(constant.REQUEST_TRACE_ID),
		)
	}
}

func LogResponse(trace string, method string, path string, status int, response interface{}) {
	log.Debug("%s: %s to endpoint %s with traceid %s ended with status %s and response \r\n $v",
		time.Now().Format(time.RFC3339),
		method,
		path,
		trace,
		status,
		response,
	)
}
