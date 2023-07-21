package controllers

import (
	"net/http"

	"github.com/bancodobrasil/featws-resolver-bridge/dtos"
	responses "github.com/bancodobrasil/featws-resolver-bridge/responses/v1"
	"github.com/bancodobrasil/featws-resolver-bridge/services"
	"github.com/gin-gonic/gin"
)

// HealthAllHandler returns a Gin handler function that checks the health of all services and returns the
// result in JSON format.
func HealthAllHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, services.CheckHealthAll())

	}
}

// HealthLiveHandler returns a Gin handler function that checks the health status of a service and returns
// it as a JSON response.
func HealthLiveHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, services.CheckHealthLive())
	}

}

// HealthReadyHandler checks if the application is ready and returns an appropriate response.
func HealthReadyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := services.FetchFromFile()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Application is not ready!!!")
			return
		}
		readyStts := services.CheckHealthReady()
		readySttsFinal := responses.NewReady(readyStts)
		if readySttsFinal.Status != "OK" {
			c.JSON(http.StatusInternalServerError, readySttsFinal)
			return
		}
		c.JSON(http.StatusOK, readySttsFinal)
	}

}

// HealthReadyServiceHandler handles requests to check the health readiness of a service and returns a response
// indicating whether the service is ready or not.
func HealthReadyServiceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		service := c.Param("service")
		readyStts := services.CheckHealthReady()
		response := dtos.Service{
			Service: service,
			Status:  readyStts.Services[service].Status,
		}
		responses := responses.NewService(response)
		if responses.Status != "200 OK" {
			responses.Status += " Service is not ready!!!"
			c.JSON(http.StatusInternalServerError, responses)
			return
		}
		responses.Status += " Service is ready!!!"
		c.JSON(http.StatusOK, responses)
	}
}
