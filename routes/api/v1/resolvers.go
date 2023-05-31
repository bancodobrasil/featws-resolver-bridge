package v1

import (
	v1 "github.com/bancodobrasil/featws-resolver-bridge/controllers/v1"
	"github.com/gin-gonic/gin"
)

// resolversRouter sets up a GET route for a specific endpoint in a Gin router group.
func resolversRouter(router *gin.RouterGroup) {
	router.GET("/", v1.ResolversHandler)
}
