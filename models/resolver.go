package models

// Resolver ...
type Resolver struct {
	Name    string                 `json:"name,omitempty"`
	Type    string                 `json:"type,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}
