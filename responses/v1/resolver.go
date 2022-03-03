package v1

// Resolver contais all output of resolver execution
type Resolver struct {
	Context map[string]interface{} `json:"context"`
	Errors  map[string]interface{} `json:"errors"`
}
