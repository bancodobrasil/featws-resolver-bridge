package services

import (
	"context"

	"github.com/bancodobrasil/featws-resolver-bridge/models"
	"github.com/bancodobrasil/featws-resolver-bridge/repository"
)

// CreateResolver ...
func CreateResolver(ctx context.Context, resolver *models.Resolver) error {

	//TODO verifica unicidade do nome

	err := repository.GetResolversRepository().Create(ctx, resolver)
	if err != nil {
		return err
	}

	return nil
}

// func GetAResolver() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		resolverId := c.Param("resolverId")
// 		var resolver models.Resolver
// 		defer cancel()

// 		objId, _ := primitive.ObjectIDFromHex(resolverId)

// 		err := resolversCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&resolver)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.ResolverResponse{
// 				Status:  http.StatusInternalServerError,
// 				Message: "error",
// 				Data:    map[string]interface{}{"data": err.Error()},
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, responses.ResolverResponse{
// 			Status:  http.StatusOK,
// 			Message: "success",
// 			Data:    map[string]interface{}{"data": resolver},
// 		})
// 	}
// }

// func EditAResolver() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		resolverId := c.Param("resolverId")
// 		var resolver models.Resolver
// 		defer cancel()
// 		objId, _ := primitive.ObjectIDFromHex(resolverId)

// 		//validate the request body
// 		if err := c.BindJSON(&resolver); err != nil {
// 			c.JSON(http.StatusBadRequest, responses.ResolverResponse{
// 				Status:  http.StatusBadRequest,
// 				Message: "error",
// 				Data:    map[string]interface{}{"data": err.Error()},
// 			})
// 			return
// 		}

// 		//use the validator library to validate required fields
// 		if validationErr := validate.Struct(&resolver); validationErr != nil {
// 			c.JSON(http.StatusBadRequest, responses.ResolverResponse{
// 				Status:  http.StatusBadRequest,
// 				Message: "error",
// 				Data:    map[string]interface{}{"data": validationErr.Error()},
// 			})
// 			return
// 		}

// 		update := bson.M{"title": resolver.Title, "author": resolver.Author, "price": resolver.Price}
// 		result, err := resolversCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.ResolverResponse{
// 				Status:  http.StatusInternalServerError,
// 				Message: "error",
// 				Data:    map[string]interface{}{"data": err.Error()},
// 			})
// 			return
// 		}

// 		//get updated resolver datails
// 		var updatedResolver models.Resolver
// 		if result.MatchedCount == 1 {
// 			err := resolversCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedResolver)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, responses.ResolverResponse{
// 					Status:  http.StatusInternalServerError,
// 					Message: "error",
// 					Data:    map[string]interface{}{"data": err.Error()},
// 				})
// 				return
// 			}

// 			c.JSON(http.StatusOK, responses.ResolverResponse{
// 				Status:  http.StatusOK,
// 				Message: "success",
// 				Data:    map[string]interface{}{"data": updatedResolver},
// 			})
// 		}
// 	}
// }

// func DeleteAResolver() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		resolverId := c.Param("resolverId")
// 		defer cancel()

// 		objId, _ := primitive.ObjectIDFromHex(resolverId)

// 		result, err := resolversCollection.DeleteOne(ctx, bson.M{"_id": objId})
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.ResolverResponse{
// 				Status:  http.StatusInternalServerError,
// 				Message: "error",
// 				Data:    map[string]interface{}{"data": err.Error()},
// 			})
// 			return
// 		}

// 		if result.DeletedCount < 1 {
// 			c.JSON(http.StatusNotFound, responses.ResolverResponse{
// 				Status:  http.StatusNotFound,
// 				Message: "error",
// 				Data:    map[string]interface{}{"data": "Resolver with specified ID not found!"}},
// 			)
// 			return
// 		}

// 		c.JSON(http.StatusNoContent, responses.ResolverResponse{
// 			Status:  http.StatusOK,
// 			Message: "success",
// 			Data:    map[string]interface{}{"data": "Resolver successfully deleted!"}},
// 		)

// 	}
// }

// func GetAllResolvers() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		var resolvers []models.Resolver
// 		defer cancel()

// 		results, err := resolversCollection.Find(ctx, bson.M{})

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.ResolverResponse{
// 				Status:  http.StatusInternalServerError,
// 				Message: "error",
// 				Data:    map[string]interface{}{"data": err.Error()},
// 			})
// 			return
// 		}

// 		// reading from the db in an optimal way
// 		defer results.Close(ctx)
// 		for results.Next(ctx) {
// 			var singleResolver models.Resolver
// 			if err = results.Decode(&singleResolver); err != nil {
// 				c.JSON(http.StatusInternalServerError, responses.ResolverResponse{
// 					Status:  http.StatusInternalServerError,
// 					Message: "error",
// 					Data:    map[string]interface{}{"data": err.Error()},
// 				})
// 				return
// 			}

// 			resolvers = append(resolvers, singleResolver)
// 		}

// 		c.JSON(http.StatusOK, responses.ResolverResponse{
// 			Status:  http.StatusOK,
// 			Message: "success",
// 			Data:    map[string]interface{}{"data": resolvers},
// 		})
// 	}
// }
