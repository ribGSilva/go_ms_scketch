package config

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("%s: %s to endpoint %s", time.Now().Format(time.RFC3339), ctx.Request.Method, ctx.Request.RequestURI)
	}
}
