package utils

import (
)

func SetDefault(value, condition, def string) string {
	if value == condition {
		return def
	}
	return value
}

func IfOrNot(truth bool, a, b interface{}) interface{} {
	if truth {
		return a
	}
	return b
}