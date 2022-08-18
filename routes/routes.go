package routes

import (
	"github.com/bancodobrasil/featws-resolver-bridge/config"
	"github.com/bancodobrasil/featws-resolver-bridge/docs"
	"github.com/bancodobrasil/featws-resolver-bridge/routes/api"
	"github.com/bancodobrasil/featws-resolver-bridge/routes/health"
	telemetry "github.com/bancodobrasil/gin-telemetry"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetupRoutes define all routes
func SetupRoutes(router *gin.Engine) {

	cfg := config.GetConfig()
	docs.SwaggerInfo.Host = cfg.ExternalHost

	homeRouter(router.Group("/"))
	health.Router(router.Group("/health"))

	// setup swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

// APIRoutes define all api routes
func APIRoutes(router *gin.Engine) {
	// inject middleware
	group := router.Group("/api")
	group.Use(telemetry.Middleware("featws-resolver-bridge"))
	api.Router(group)
}
