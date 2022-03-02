package routes

import (
	"github.com/bancodobrasil/featws-resolver-adapter-go/routes/api"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	homeRouter(router.Group("/"))
	api.ApiRouter(router.Group("/api"))
}
