package v1

// Resolver contains all input for resolver execution
// type Resolver struct {
// 	Context map[string]interface{} `json:"context"`
// 	Load    []string               `json:"load"`
// }
type Resolver struct {
	ID      string                 `json:"id,omitempty" mapstructure:"id,omitempty"`
	Name    string                 `json:"name,omitempty" validate:"required" mapstructure:"name"`
	Type    string                 `json:"type,omitempty" validate:"required" mapstructure:"type"`
	Options map[string]interface{} `json:"options,omitempty" validate:"required" mapstructure:"options"`
}
