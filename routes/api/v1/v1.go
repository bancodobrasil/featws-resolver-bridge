package v1

import (
	"github.com/gin-gonic/gin"
)

// Router define routes the API V1
func Router(router *gin.RouterGroup) {
	resolversRouter(router.Group("/resolvers"))
	rpcRouter(router.Group("/"))
}
