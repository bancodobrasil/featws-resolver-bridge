package controllers

import (
	"net/http"

	"github.com/bancodobrasil/featws-resolver-bridge/services"
	"github.com/gin-gonic/gin"
)

// HealthReadyHandler ...
func HealthReadyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := services.FetchFromFile()
		if err != nil {
			c.String(http.StatusInternalServerError, "Application is not ready!!!")
			return
		}
		c.String(http.StatusOK, "Application is Ready!!!")
	}

}
