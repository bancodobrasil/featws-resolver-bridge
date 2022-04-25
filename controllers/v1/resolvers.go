package v1

import (
	"context"
	"net/http"
	"time"

	responses "github.com/bancodobrasil/featws-resolver-bridge/responses/v1"
	"github.com/bancodobrasil/featws-resolver-bridge/services"
	"github.com/gin-gonic/gin"
)

// ResolversHandler godoc
// @Summary 		Create Resolvers
// @Description Create Resolvers description
// @Tags 				resolvers
// @Accept  		json
// @Produce  		json
// @Success 		200 {object} string "ok"
// @Failure 		400,404 {object} string
// @Failure 		500 {object} string
// @Failure 		default {object} string
// @Router 			/resolvers/ [get]
// @Router 			/resolve/{resolver} [post]
// @Router 			/resolve [post]
// @Router 			/load [get]
func ResolversHandler(c *gin.Context) {

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	entities, err := services.FetchResolvers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{
			Error: err.Error(),
		})
		return
	}

	var response = make([]responses.Resolver, len(entities))

	for index, entity := range entities {
		response[index] = responses.NewResolver(entity)
	}

	c.JSON(http.StatusOK, response)
}
