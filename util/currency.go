package util

import "strings"

var currencies = []string{"USD", "EUR"}

func IsValidCurrency(currencyString string) bool {

	for _, curency := range currencies {
		if curency == strings.Trim(currencyString, " ") {
			return true
		}
	}

	return false
}
