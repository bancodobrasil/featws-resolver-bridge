package models

import (
	v1 "github.com/bancodobrasil/featws-resolver-bridge/payloads/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Resolver ...
type Resolver struct {
	ID      primitive.ObjectID     `bson:"_id,omitempty"`
	Name    string                 `bson:"name,omitempty"`
	Type    string                 `bson:"type,omitempty"`
	Options map[string]interface{} `bson:"options,omitempty"`
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
		Type:    payload.Type,
		Options: payload.Options,
	}
	return
}
