package utils

import (
	"path/filepath"
)

func JoinPath(paths...string) string {
	return filepath.Join(paths...)
}