package utils

import (
	"strings"
)

func ReplaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

func ReplaceAll(s string, oldnew ...string) string {
	repl := strings.NewReplacer(oldnew...)
	return repl.Replace(s)
}
