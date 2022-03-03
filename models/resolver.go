package models

type Resolver struct {
	Id      interface{}       `json:"id,omitempty"`
	Name    string            `json:"name,omitempty" validate:"required"`
	Url     string            `json:"url,omitempty" validate:"required"`
	Headers map[string]string `json:"headers,omitempty"`
}
