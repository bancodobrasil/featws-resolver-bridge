package v1

import "github.com/bancodobrasil/featws-resolver-bridge/models"

// Resolver contais all output of resolver execution
//
//	type Resolver struct {
//		Context map[string]interface{} `json:"context"`
//		Errors  map[string]interface{} `json:"errors"`
//	}

// Resolver has three fields: `Name`, `Type`, and `Version`, each of which is a string and
// can be represented in JSON format with corresponding keys.
//
// Property:
//   - {string} Name: is a string property that represents the name of a resolver. It is tagged with `json:"name,omitempty"` which means that it will be included in the JSON representation of the struct only if it has a non-zero value.
//   - {string} Type - The "Type" property is a string field that represents the type of the resolver. It could be a DNS resolver, a proxy resolver, or any other type of resolver.
//   - {string} Version - The "Version" property is a string that represents the version number of a software component or module. It is used to track changes and updates to the software over time.
type Resolver struct {
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
}

// NewResolver creates a new Resolver object based on a given entity.
func NewResolver(entity models.Resolver) Resolver {
	return Resolver{
		Name:    entity.Name,
		Type:    entity.Type,
		Version: entity.Version,
	}
}
