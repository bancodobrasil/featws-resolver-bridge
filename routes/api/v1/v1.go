package v1

import (
	"github.com/gin-gonic/gin"
)

func V1Router(router *gin.RouterGroup) {
	resolverRouter(router.Group("/resolver"))
}
