package grpcutil

import (
	"path"
	"strings"
)

// CanonicalMethodPath ensures the method path starts with a single slash.
func CanonicalMethodPath(p string) string {
	p = path.Clean("/" + p)
	if !strings.HasPrefix(p, "/") {
		return "/" + p
	}
	return p
}