package routes

import (
	"github.com/bancodobrasil/featws-resolver-adapter-go/controllers"
	"github.com/gin-gonic/gin"
)

func homeRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.HomeHandler)
}
