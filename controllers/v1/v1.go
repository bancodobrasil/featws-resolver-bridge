package v1

import "github.com/go-playground/validator/v10"

// `var validate = validator.New()` is creating a new instance of the `validator` struct from the
// `github.com/go-playground/validator/v10` package and assigning it to the variable `validate`. This
// allows the `validate` variable to be used to validate data using the validation rules and tags
// provided by the `validator` package.
var validate = validator.New()
