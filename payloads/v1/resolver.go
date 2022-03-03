package v1

// Resolver contains all input for resolver execution
type Resolver struct {
	Context map[string]interface{} `json:"context"`
	Load    []string               `json:"load"`
}
