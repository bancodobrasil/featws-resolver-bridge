package v1

import "github.com/bancodobrasil/featws-resolver-bridge/dtos"

// Resolve contais all output of resolver execution
type Resolve struct {
	Context map[string]interface{} `json:"context"`
	Errors  map[string]interface{} `json:"errors"`
}

// NewResolve ...
func NewResolve(dto dtos.ResolveContext) Resolve {
	return Resolve{
		Context: dto.Context,
		Errors:  dto.Errors,
	}
}
