package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func HomeHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.String(http.StatusOK, "FeatWS Resolver Works!!!")
// 	}

// }

// HomeHandler returns a string indicating that the FeatWS Resolver Bridge works.
func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "FeatWS Resolver Bridge Works!!!")
}
