package v1

import (
	"net/http"

	"github.com/bancodobrasil/featws-resolver-bridge/services"
	"github.com/gin-gonic/gin"
)

// ResolverHandler ...
func ResolverHandler(c *gin.Context) {
	var input services.ResolveInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resolverOutput := services.Resolve(input)
	c.JSON(http.StatusOK, resolverOutput)
}
