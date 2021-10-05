package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func CompareVersion(src, toCompare string) bool {
	if toCompare == "" {
		return false
	}

	exp, _ := regexp.Compile(`-(.*)`)
	src = exp.ReplaceAllString(src, "")
	toCompare = exp.ReplaceAllString(toCompare, "")

	srcs := strings.Split(src, "v")
	srcArr := strings.Split(srcs[1], ".")
	op := ">"
	srcs[0] = strings.TrimSpace(srcs[0])
	if Contains([]string{">=", "<=", "=", ">", "<"}, srcs[0]) {
		op = srcs[0]
	}

	toCompare = strings.ReplaceAll(toCompare, "v", "")

	if op == "=" {
		return srcs[1] == toCompare
	}

	if srcs[1] == toCompare && (op == "<=" || op == ">=") {
		return true
	}

	toCompareArr := strings.Split(strings.ReplaceAll(toCompare, "v", ""), ".")
	for i := 0; i < len(srcArr); i++ {
		v, err := strconv.Atoi(srcArr[i])
		if err != nil {
			return false
		}
		vv, err := strconv.Atoi(toCompareArr[i])
		if err != nil {
			return false
		}
		switch op {
		case ">", ">=":
			if v < vv {
				return true
			} else if v > vv {
				return false
			} else {
				continue
			}
		case "<", "<=":
			if v > vv {
				return true
			} else if v < vv {
				return false
			} else {
				continue
			}
		}
	}

	return false
}
