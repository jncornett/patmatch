package patmatch

import (
	"regexp"
	"strings"
)

// Interpolate substitutes variables in s for strings in values.
// If a variable in s does not exist in values, Interpolate will substitute the
// empty string.
func Interpolate(s string, values map[string]string) string {
	return variableRe.ReplaceAllStringFunc(s, func(m string) string {
		key := strings.TrimRight(strings.TrimLeft(m, "["), "]")
		return values[key]
	})
}

// InterpolatePartial substitutes variables in s for strings in values.
// If a variable in s does not exist in values, InterpolatePartial will leave it alone.
func InterpolatePartial(s string, values map[string]string) string {
	return variableRe.ReplaceAllStringFunc(s, func(m string) string {
		key := strings.TrimRight(strings.TrimLeft(m, "["), "]")
		val, exists := values[key]
		if !exists {
			return m
		}
		return val
	})
}

var variableRe = regexp.MustCompile(`\[([\w_-]+)\]`)
