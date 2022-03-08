package v1

import "github.com/bancodobrasil/featws-resolver-bridge/models"

// Resolver ...
type Resolver struct {
	ID      string                 `json:"id,omitempty"`
	Name    string                 `json:"name,omitempty"`
	Type    string                 `json:"type,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}

// NewResolver ...
func NewResolver(entity models.Resolver) Resolver {
	return Resolver{
		ID:      entity.ID.Hex(),
		Name:    entity.Name,
		Type:    entity.Type,
		Options: entity.Options,
	}
}
