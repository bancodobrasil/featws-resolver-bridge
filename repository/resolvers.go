package repository

import (
	"context"

	"github.com/bancodobrasil/featws-resolver-bridge/database"
	"github.com/bancodobrasil/featws-resolver-bridge/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// Resolvers ...
type Resolvers struct {
	collection *mongo.Collection
}

var instance = Resolvers{}

// GetResolversRepository ...
func GetResolversRepository() Resolvers {
	if instance.collection == nil {
		instance.collection = database.GetCollection("resolvers")
	}

	return instance
}

// Create ...
func (s Resolvers) Create(ctx context.Context, resolver *models.Resolver) error {

	result, err := s.collection.InsertOne(ctx, resolver)
	if err != nil {
		return err
	}

	resolver.ID = result.InsertedID

	return nil
}
