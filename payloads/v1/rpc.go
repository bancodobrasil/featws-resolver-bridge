package v1

// Resolve contains all input for resolver execution
// Property:
//   - {string} Resolver: is a string that specifies the name of the resolver function that will be used to resolve the data for a particular field in a GraphQL query. The resolver function is responsible for fetching the data from a data source and returning it to the client.
//   - Context: is a map of key-value pairs that provide additional context or information to the resolver function. The keys are strings and the values can be of any type. This context can be used by the resolver function to make decisions or perform actions based on the provided information.
//   - {[]string} Load: is an array of strings that specifies the modules or files that need to be loaded before the resolver can be executed. These can be local files or external modules that are required for the resolver to function properly.
//   - Options: : is a map of additional configuration options for the resolver. These options can vary depending on the specific resolver being used. For example, if the resolver is a database connection, the options might include the database host, port, username, and password. The Options field in the Resolve struct stores these options as a map of key-value pairs, where the key is a string representing the name of the option and the value can be of any type.
type Resolve struct {
	Resolver string                 `json:"resolver"`
	Context  map[string]interface{} `json:"context"`
	Load     []string               `json:"load"`
	Options  map[string]interface{} `json:"options"`
}
