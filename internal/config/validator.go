package config

import "github.com/go-playground/validator/v10"

// StructValidator wraps the playground validator to match Fiber's interface
type StructValidator struct {
	Validator *validator.Validate
}

// Validate implements Fiber's StructValidator interface
func (v *StructValidator) Validate(out any) error {
	return v.Validator.Struct(out)
}

// NewValidator initializes and returns the custom validator
func NewValidator() *StructValidator {
	return &StructValidator{
		Validator: validator.New(),
	}
}
