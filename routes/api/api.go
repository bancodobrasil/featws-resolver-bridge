package api

import (
	v1 "github.com/bancodobrasil/featws-resolver-adapter-go/routes/api/v1"
	"github.com/gin-gonic/gin"
)

func ApiRouter(router *gin.RouterGroup) {
	v1.V1Router(router.Group("/v1"))
}
