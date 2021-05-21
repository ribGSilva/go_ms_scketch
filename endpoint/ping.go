package endpoint

import "github.com/gin-gonic/gin"

func InjectPingEndpoints(router *gin.Engine) {
	router.GET("/ping", getPing)
}

func getPing(ctx *gin.Context) {

}
