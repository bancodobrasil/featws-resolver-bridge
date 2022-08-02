package routes

import (
	"github.com/bancodobrasil/featws-resolver-bridge/routes/api"
	"github.com/bancodobrasil/featws-resolver-bridge/routes/health"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetupRoutes define all routes
func SetupRoutes(router *gin.Engine) {

	homeRouter(router.Group("/"))
	api.Router(router.Group("/api"))
	health.Router(router.Group("/health"))

	// setup swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
