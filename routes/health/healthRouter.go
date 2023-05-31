package health

import (
	"github.com/bancodobrasil/featws-resolver-bridge/controllers"
	"github.com/gin-gonic/gin"
)

// Router sets up the routing for various health check endpoints in a Go web application using
// the Gin framework.
func Router(router *gin.RouterGroup) {
	router.GET("/", controllers.HealthAllHandler())
	router.GET("/live", controllers.HealthLiveHandler())
	router.GET("/ready", controllers.HealthReadyHandler())
	router.GET("/ready/:service", controllers.HealthReadyServiceHandler())
}
