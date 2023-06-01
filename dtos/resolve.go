package dtos

import v1 "github.com/bancodobrasil/featws-resolver-bridge/payloads/v1"

// ResolveContext type contains maps for context, load, options, and errors in Go.
//
// Property:
//   - Context: is a map that stores key-value pairs of data that can be accessed and used by the resolver functions. The keys in the map represent the names of the data, while the values represent the actual data. This data can be used to provide additional context to the resolver functions, such as user information, configuration settings, or any other relevant data.
//   - {[]string} Load - The `Load` property is a slice of strings that represents the list of files or resources that need to be loaded before resolving the context. These files or resources can contain additional data or configuration that is required to properly resolve the context.
//   - Options: is a map that contains various options that can be used during the resolution process. These options can be set by the user to customize the behavior of the resolver. The specific options available and their meanings depend on the implementation of the resolver.
//   - Errors - The `Errors` property is a map that stores any errors encountered during the resolution process. The keys of the map represent the names of the properties that had errors, and the values represent the error messages. This allows for easy tracking and handling of errors during the resolution process.
type ResolveContext struct {
	Context map[string]interface{}
	Load    []string
	Options map[string]interface{}
	Errors  map[string]interface{}
}

// NewResolveV1 creates a new ResolveContext object with the given payload and initializes its Errors
// field with an empty map.
func NewResolveV1(payload v1.Resolve) ResolveContext {
	return ResolveContext{
		Context: payload.Context,
		Load:    payload.Load,
		Options: payload.Options,
		Errors:  make(map[string]interface{}),
	}
}
