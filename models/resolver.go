package models

// Resolver ...
type Resolver struct {
	Name    string                 `json:"name,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
	Type    string                 `json:"type,omitempty"`
	Version string                 `json:"version,omitempty"`
}
