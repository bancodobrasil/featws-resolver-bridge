package v1

import "github.com/bancodobrasil/featws-resolver-bridge/models"

// Resolver contais all output of resolver execution
// type Resolver struct {
// 	Context map[string]interface{} `json:"context"`
// 	Errors  map[string]interface{} `json:"errors"`
// }
type Resolver struct {
	ID      interface{}       `json:"id,omitempty"`
	Name    string            `json:"name,omitempty"`
	URL     string            `json:"url,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

// NewResolver ...
func NewResolver(entity models.Resolver) Resolver {
	return Resolver{
		ID:      entity.ID,
		Name:    entity.Name,
		URL:     entity.URL,
		Headers: entity.Headers,
	}
}
