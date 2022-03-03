package v1

// ResolveHandler ...
// func ResolveHandler(c *gin.Context) {
// 	var input payloads.Resolver
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	resolverOutput := services.Resolve(input)
// 	c.JSON(http.StatusOK, resolverOutput)
// }
