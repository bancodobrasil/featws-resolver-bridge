package v1

import (
	v1 "github.com/bancodobrasil/featws-resolver-bridge/controllers/v1"
	"github.com/gin-gonic/gin"
)

func rpcRouter(router *gin.RouterGroup) {
	router.POST("/resolve", v1.ResolveHandler)
}
