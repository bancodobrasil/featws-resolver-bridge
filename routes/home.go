package routes

import (
	"github.com/bancodobrasil/featws-resolver-bridge/controllers"
	"github.com/gin-gonic/gin"
)

func homeRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.HomeHandler)
}
