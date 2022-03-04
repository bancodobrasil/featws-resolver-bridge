package models

import v1 "github.com/bancodobrasil/featws-resolver-bridge/payloads/v1"

// Resolver ...
type Resolver struct {
	ID      interface{}
	Name    string
	URL     string
	Headers map[string]string
}

// NewResolverV1 ...
func NewResolverV1(payload v1.Resolver) *Resolver {
	return &Resolver{
		ID:      payload.ID,
		Name:    payload.Name,
		URL:     payload.URL,
		Headers: payload.Headers,
	}
}
