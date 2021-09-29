package utils

import (
	"encoding/json"
	"strconv"
)

func ParseJSON(a interface{}) string {
	if a == nil {
		return ""
	}
	b, _ := json.Marshal(a)
	return string(b)
}

func ParseBool(s string) bool {
	b1, _ := strconv.ParseBool(s)
	return b1
}

func ParseFloat32(f string) float32 {
	s, _ := strconv.ParseFloat(f, 32)
	return float32(s)
}
