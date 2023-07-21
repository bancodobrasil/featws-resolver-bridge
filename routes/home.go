package routes

import (
	"github.com/bancodobrasil/featws-resolver-bridge/controllers"
	"github.com/gin-gonic/gin"
)

// homeRouter sets up a GET route for the home page using the Gin web framework in Go.
func homeRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.HomeHandler)
}
