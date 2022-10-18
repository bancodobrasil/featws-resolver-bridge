package v1

import (
	"github.com/bancodobrasil/featws-resolver-bridge/middlewares"
	"github.com/gin-gonic/gin"
)

// Router define routes the API V1
func Router(router *gin.RouterGroup) {
	router.Use(middlewares.VerifyAPIKey())
	resolversRouter(router.Group("/resolvers"))
	rpcRouter(router.Group("/"))
}
