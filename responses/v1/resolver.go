package v1

import "github.com/bancodobrasil/featws-resolver-bridge/models"

// Resolver contais all output of resolver execution
// type Resolver struct {
// 	Context map[string]interface{} `json:"context"`
// 	Errors  map[string]interface{} `json:"errors"`
// }
type Resolver struct {
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
}

// NewResolver ...
func NewResolver(entity models.Resolver) Resolver {
	return Resolver{
		Name:    entity.Name,
		Type:    entity.Type,
		Version: entity.Version,
	}
}
