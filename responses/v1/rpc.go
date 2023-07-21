package v1

import "github.com/bancodobrasil/featws-resolver-bridge/dtos"

// Resolve contais all output of resolver execution. Resolve has two fields, `Context` and `Errors`, both of which are maps with string keys
// and interface values, and is used for JSON serialization with specific field names.
//
// Property:
//
//   - Context: is a map that stores key-value pairs of contextual information related to a particular operation or request. This information can be used to provide additional context to the operation or request, such as user information, device information, or any other relevant data.
//   - Errors: is a map that contains any errors or issues encountered during the resolution process. The keys in the map represent the names of the fields or properties that had issues, and the values contain information about the specific error or issue that occurred. This property is commonly used in GraphQL to provide
type Resolve struct {
	Context map[string]interface{} `json:"context"`
	Errors  map[string]interface{} `json:"errors"`
}

// NewResolve creates a new Resolve object using data from a ResolveContext DTO.
func NewResolve(dto dtos.ResolveContext) Resolve {
	return Resolve{
		Context: dto.Context,
		Errors:  dto.Errors,
	}
}
