package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/bancodobrasil/featws-resolver-bridge/dtos"
	payloads "github.com/bancodobrasil/featws-resolver-bridge/payloads/v1"
	responses "github.com/bancodobrasil/featws-resolver-bridge/responses/v1"
	"github.com/bancodobrasil/featws-resolver-bridge/services"
	"github.com/gin-gonic/gin"
)

// ResolveHandler ...
func ResolveHandler(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var input payloads.Resolve
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{
			Error: err.Error(),
		})
		return
	}

	resolveContext := dtos.NewResolveV1(input)

	resolver := input.Resolver

	resolverName, exists := c.Params.Get("resolver")

	if exists {
		resolver = resolverName
	}

	err := services.Resolve(ctx, resolver, &resolveContext)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{
			Error: err.Error(),
		})
		return
	}

	resolverOutput := responses.NewResolve(resolveContext)

	c.JSON(http.StatusOK, resolverOutput)
}

// LoadHandler ...
func LoadHandler(c *gin.Context) {

	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := services.Load()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{
			Error: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
