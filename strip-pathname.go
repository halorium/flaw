package flaw

import (
	"regexp"
)

// stripPathname shortens the path to everything after the repository root
// for simplicity purposes
func stripPathname(pathname string, repoRoot string) string {

	rgx := regexp.MustCompile(repoRoot + "/(.+)")

	pathMatches := rgx.FindAllStringSubmatch(pathname, -1)

	if pathMatches == nil {
		return pathname
	}

	return pathMatches[0][1]
}
