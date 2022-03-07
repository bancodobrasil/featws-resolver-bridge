package repository

import (
	"context"

	"github.com/bancodobrasil/featws-resolver-bridge/database"
	"github.com/bancodobrasil/featws-resolver-bridge/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (r Resolvers) Create(ctx context.Context, resolver *models.Resolver) error {

	result, err := r.collection.InsertOne(ctx, resolver)
	if err != nil {
		return err
	}

	resolver.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}

// Find ...
func (r Resolvers) Find(ctx context.Context) (list []models.Resolver, err error) {
	results, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var resolver models.Resolver
		if err = results.Decode(&resolver); err != nil {
			return
		}

		list = append(list, resolver)
	}

	return
}

// Get ...
func (r Resolvers) Get(ctx context.Context, id string) (resolver *models.Resolver, err error) {

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	result := r.collection.FindOne(ctx, bson.M{"_id": oid})

	err = result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return
	}

	result.Decode(&resolver)

	return
}

func (r Resolvers) Update(ctx context.Context, entity models.Resolver) (updated *models.Resolver, err error) {

	update := bson.M{"name": entity.Name, "url": entity.URL, "headers": entity.Headers}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": entity.ID}, bson.M{"$set": update})

	if err != nil {
		return
	}

	updated, err = r.Get(ctx, entity.ID.Hex())
	if err != nil {
		return
	}

	return
}
