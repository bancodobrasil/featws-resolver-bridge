package models

// Resolver contains information about a resolver, including its name, options, type, and version.
//
// Property:
//   - {string} Name: is a string property that represents the name of the resolver. It is marked as "omitempty" in the JSON tag, which means that if the value of this property is empty, it will be omitted from the JSON output.
//   - Options: is a map that can contain any additional configuration options for the resolver. The keys of the map are strings that represent the name of the option, and the values can be of any type. This property is marked as omitempty, which means that if the map is empty, it will be omitted when the struct is serialized to JSON. The Options field is part of the Resolver struct, which also includes other properties such as Name, Type, and Version.
//   - {string} Type: is a string that represents the type of the resolver. It could be a DNS resolver, a proxy resolver, or any other type of resolver.
//   - Version: is a string that represents the version of the resolver. It is used to keep track of changes and updates made to the resolver over time.
type Resolver struct {
	Name    string                 `json:"name,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
	Type    string                 `json:"type,omitempty"`
	Version string                 `json:"version,omitempty"`
}
