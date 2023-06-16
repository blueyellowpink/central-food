package validator

import (
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
)

func BindValidators() error {
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validEmail", validEmail)
		v.RegisterValidation("validPassword", validPassword)
	}
    return nil
}

var validEmail validator.Func = func(fl validator.FieldLevel) bool {
    _, ok := fl.Field().Interface().(string)
	if ok {
        return true
	}
	return false
}

var validPassword validator.Func = func(fl validator.FieldLevel) bool {
    _, ok := fl.Field().Interface().(string)
	if ok {
        return true
	}
	return false
}
