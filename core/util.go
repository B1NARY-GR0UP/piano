package core

import (
	"bytes"
	"strings"

	"github.com/B1NARY-GR0UP/inquisitor/core"
)

// M music is used to simplified code
type M map[string]any

// validateRoute return true if the route is valid
func validateRoute(method, path string, handlers HandlersChain) bool {
	if method == "" {
		core.Info("[PIANO] HTTP method can not be empty")
		return false
	}
	if path[0] != '/' {
		core.Info("[PIANO] path must start with '/'")
		return false
	}
	if len(handlers) < 1 {
		core.Info("[PIANO] there must be at least one handler")
		return false
	}
	return true
}

// validatePath check URL path if it's valid
func validatePath(path string) {
	if path == nullString {
		panic("path is empty")
	}
	if path[0] != charSlash {
		panic("path must begin with '/'")
	}
	for _, c := range []byte(path) {
		// TODO: enrich logic
		switch c {
		case charColon:

		case charStar:

		}
	}
}

// splitPath split the URL path into fragments
// e.g. /binary/lorain => [binary lorain]
func splitPath(path string) []string {
	return strings.Split(path, strSlash)[1:]
}

// calculateParam calculate the count of special fragments in a path
func calculateParam(path string) uint16 {
	return uint16(bytes.Count([]byte(path), []byte(strColon)) + bytes.Count([]byte(path), []byte(strStar)))
}
