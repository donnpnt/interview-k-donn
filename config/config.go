package config

import (
	"os"
	"strings"
)

func ResolveBasePath(path ...string) string {
	p, _ := os.Getwd()

	if len(path) > 0 {
		p = p + "/" + strings.TrimLeft(path[0], "/")
	}

	return p
}
