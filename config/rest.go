package config

import (
	"fmt"
	"gabrielrsilva/go-ms-scketch/constant"
	"gabrielrsilva/go-ms-scketch/object"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ConfigRouter(router *gin.Engine) {
	router.Use(newAuthorizationHandler())
	router.Use(newRequestTraceIdHandler())
}

func newAuthorizationHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader(constant.AUTHORIZATION)
		if authorization == "" {
			err := constant.NotAuthorizedErrorTemplate
			ctx.AbortWithStatusJSON(http.StatusBadRequest, object.NewErrorResponse(
				err.Code,
				err.Message,
				fmt.Sprintf(err.DetailTemplate, "Missing authorization"),
			))
		}
	}
}

func newRequestTraceIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trace := ctx.GetHeader(constant.REQUEST_TRACE_ID)
		if trace == "" {
			id, _ := uuid.NewRandom()
			ctx.Request.Header.Add(constant.REQUEST_TRACE_ID, id.String())
		}
	}
}
