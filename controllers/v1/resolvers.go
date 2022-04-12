package v1

import (
	"context"
	"net/http"
	"time"

	responses "github.com/bancodobrasil/featws-resolver-bridge/responses/v1"
	"github.com/bancodobrasil/featws-resolver-bridge/services"
	"github.com/gin-gonic/gin"
)

// ResolversHandler ...
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
