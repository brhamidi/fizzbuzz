package http

import (
	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
// https://github.com/go-playground/validator/blob/master/_examples/simple/main.go#L27
var validate *validator.Validate

func init() {
	validate = validator.New()
}
