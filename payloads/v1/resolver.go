package v1

// Resolver contains all input for resolver execution
// type Resolver struct {
// 	Context map[string]interface{} `json:"context"`
// 	Load    []string               `json:"load"`
// }
type Resolver struct {
	ID      string            `json:"id,omitempty"`
	Name    string            `json:"name,omitempty" validate:"required"`
	URL     string            `json:"url,omitempty" validate:"required"`
	Headers map[string]string `json:"headers,omitempty"`
}