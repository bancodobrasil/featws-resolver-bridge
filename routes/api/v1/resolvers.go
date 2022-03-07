package v1

import (
	v1 "github.com/bancodobrasil/featws-resolver-bridge/controllers/v1"
	"github.com/gin-gonic/gin"
)

func resolversRouter(router *gin.RouterGroup) {
	router.POST("/", v1.CreateResolver())
	router.GET("/", v1.GetResolvers())
	router.GET("/:id", v1.GetResolver())
	router.PUT("/:id", v1.UpdateResolver())
}
