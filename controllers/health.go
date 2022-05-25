package controllers

import (
	"net/http"

	"github.com/bancodobrasil/featws-resolver-bridge/dtos"
	responses "github.com/bancodobrasil/featws-resolver-bridge/responses/v1"
	"github.com/bancodobrasil/featws-resolver-bridge/services"
	"github.com/gin-gonic/gin"
)

// HealthAllHandler ...
func HealthAllHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, services.CheckHealthAll())

	}
}

// HealthLiveHandler ...
func HealthLiveHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, services.CheckHealthLive())
	}

}

// HealthReadyHandler ...
func HealthReadyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Test if the resolvers.json is available
		_, err := services.FetchFromFile()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Application is not ready!!!")
			return
		}

		// var code []string
		// Test the if each resolver were available
		readyStts := services.CheckHealthReady()
		readySttsFinal := responses.NewReady(readyStts)
		if readySttsFinal.Status != "OK" {
			c.JSON(http.StatusInternalServerError, readySttsFinal)
			return
		}
		c.JSON(http.StatusOK, readySttsFinal)
		// re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		// for _, ready := range readyStts {

		// 	code = re.FindAllString(ready.Status, -1)
		// 	if !reflect.DeepEqual(code, []string{"200"}) {
		// 		c.JSON(http.StatusInternalServerError, ready)
		// 	} else {
		// 	}
		// }
	}

}

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
		// if _, ok := readyStts.Services[service]; !ok {
		// 	c.JSON(http.StatusNotFound, readyStts)
		// 	return
		// }
		c.JSON(http.StatusOK, responses)
	}
}
