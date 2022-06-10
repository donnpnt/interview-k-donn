package helpers

import (
	"strings"

	"github.com/zipzoft/interview-k-donn/config"
)

func StoragePath(path ...string) string {
	p := config.ResolveBasePath("storage")
	if len(path) > 0 {
		p = p + "/" + strings.TrimLeft(path[0], "/")
	}

	return p
}

func CachePath(path ...string) string {
	p := StoragePath("cache")
	if len(path) > 0 {
		p = p + "/" + strings.TrimLeft(path[0], "/")
	}

	return p
}
