package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InjectHelloEndpoit(router *gin.Engine) {
	router.GET("/hello", getHello)
}

func getHello(ctx *gin.Context) {
	firstname := ctx.DefaultQuery("firstname", "Guest")
	lastname := ctx.Query("lastname")

	ctx.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
