package api

import (
	"github.com/Nuwan-Walisundara/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		//valid string
		return util.IsValidCurrency(currency)
	}
	return false
}
