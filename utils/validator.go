package utils

import validator "gopkg.in/go-playground/validator.v9"

type GeneralValidator struct {
	Validator *validator.Validate
}

func (v *GeneralValidator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}
