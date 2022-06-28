package dtos

import v1 "github.com/bancodobrasil/featws-resolver-bridge/payloads/v1"

// ResolveContext ...
type ResolveContext struct {
	Context map[string]interface{}
	Load    []string
	Options map[string]interface{}
	Errors  map[string]interface{}
}

// NewResolveV1 ...
func NewResolveV1(payload v1.Resolve) ResolveContext {
	return ResolveContext{
		Context: payload.Context,
		Load:    payload.Load,
		Options: payload.Options,
		Errors:  make(map[string]interface{}),
	}
}
