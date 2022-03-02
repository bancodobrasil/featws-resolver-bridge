package v1

import (
	"net/http"

	"github.com/bancodobrasil/featws-resolver-adapter-go/services"
	"github.com/gin-gonic/gin"
)

func ResolverHandler(c *gin.Context) {
	var input services.ResolveInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resolverOutput := services.Resolve(input)
	c.JSON(http.StatusOK, resolverOutput)
}
