package models

import (
	v1 "github.com/bancodobrasil/featws-resolver-bridge/payloads/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Resolver ...
type Resolver struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name,omitempty"`
	URL     string             `bson:"url,omitempty"`
	Headers map[string]string  `bson:"headers,omitempty"`
}

// NewResolverV1 ...
func NewResolverV1(payload v1.Resolver) (entity Resolver, err error) {

	id := primitive.NilObjectID

	if payload.ID != "" {
		id, err = primitive.ObjectIDFromHex(payload.ID)
		if err != nil {
			return
		}
	}

	entity = Resolver{
		ID:      id,
		Name:    payload.Name,
		URL:     payload.URL,
		Headers: payload.Headers,
	}
	return
}
