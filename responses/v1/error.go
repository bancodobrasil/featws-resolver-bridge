package v1

// Error is an error object with a flexible error field that can hold any type of data.
//
// Property
//   - Error: ia a struct is a custom error type in Go that has a single field named `Error` of type `interface{}`. This field is tagged with `json:"error"`, which means that when the struct is serialized to JSON.
type Error struct {
	Error interface{} `json:"error"`
}
