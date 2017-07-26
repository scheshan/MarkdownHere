package server

import (
	"strings"
)

func isFolder(path string) bool {
	return strings.HasSuffix(path, "/")
}
