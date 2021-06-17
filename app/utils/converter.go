package utils

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// CamelToSnakeCase func
func CamelToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// FloatNullable func
func FloatNullable(value interface{}) float64 {

	if value == nil {
		return 0
	}

	return value.(float64)
}

// StringNullable func
func StringNullable(value interface{}) string {

	if value != nil {
		if len(value.(string)) <= 0 {
			return ""
		}
	} else {
		return ""
	}

	return value.(string)
}

// IntNullable func
func IntNullable(value interface{}) int {

	if value == nil {
		return 0
	}

	return value.(int)
}

// BoolToInteger func
func BoolToInteger(value bool) int {

	if value == true {
		return 1
	}

	return 0
}

// IntNullableForIsActive func
func IntNullableForIsActive(value interface{}) int {

	if value == nil {
		return 1
	}

	return value.(int)
}
