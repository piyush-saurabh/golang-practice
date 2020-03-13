package types

import (
	vCop "gopkg.in/go-playground/validator.v9"
)

var validator *vCop.Validate

func init() {
	validator = vCop.New()
}

// Validate method validates the struct fields based on it's tag `validate`
func Validate(t interface{}) error {
	return validator.Struct(t)
}
