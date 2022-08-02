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
	log "github.com/sirupsen/logrus"
)

// ResolveHandler godoc
// @Summary 		Execute the Resolve resolutions
// @Description Resolvers description
// @Tags 				resolve
// @Accept  		json
// @Produce  		json
// @Param           resolver path string false "resolver"
// @Param           parameters body payloads.Resolve true "Parameters"
// @Success 		200 {object} string "ok"
// @Failure 		400,404 {object} string
// @Failure 		500 {object} string
// @Failure 		default {object} string
// @Router 			/resolve/{resolver} [post]
func ResolveHandler(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	var input payloads.Resolve
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Errorf("error occurs on binding JSON: %v", err)
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
		log.Errorf("error occurs on resolve the context: %v", err)
		c.JSON(http.StatusInternalServerError, responses.Error{
			Error: err.Error(),
		})
		return
	}

	resolverOutput := responses.NewResolve(resolveContext)

	c.JSON(http.StatusOK, resolverOutput)
}

// LoadHandler godoc
// @Summary 		Load Resolvers
// @Description Load Resolvers description
// @Tags 				load
// @Accept  		json
// @Produce  		json
// @Success 		200 {object} string "ok"
// @Failure 		400,404 {object} string
// @Failure 		500 {object} string
// @Failure 		default {object} string
// @Router 			/load [get]
func LoadHandler(c *gin.Context) {

	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := services.Load()
	if err != nil {
		log.Errorf("error occurs on load: %v", err)
		c.JSON(http.StatusInternalServerError, responses.Error{
			Error: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
