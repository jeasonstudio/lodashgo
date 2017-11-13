package String

import (
	"strings"
)

// EndsWith Checks if string ends with the given target string.
func EndsWith(source string, target string) bool {
	selfTarget, selfSource := strings.Split(target, ""), strings.Split(source, "")
	return strings.Join(selfSource[len(selfSource) - len(selfTarget):], "") == strings.Join(selfTarget, "")
}
