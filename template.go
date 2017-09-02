package patmatch

import "regexp"

// Template is the representation of a compiled template expression.
// A Template is safe for concurrent use by multiple goroutines.
type Template struct {
	re    *regexp.Regexp
	names []string
}

// Match returns true if the template matches the given string.
func (m *Template) Match(s string) bool {
	if m == nil {
		return false
	}
	return m.re.MatchString(s)
}

// Apply applies the template to the given string and returns all the named
// capture groups. If there was no match, Apply returns nil.
// If there was a match, even if there were no named capture groups,
// Apply will return a 1-element map with the matched expression keyed by the
// empty string.
func (m *Template) Apply(s string) map[string]string {
	if m == nil {
		return nil
	}
	result := m.re.FindStringSubmatch(s)
	if result == nil {
		return nil
	}
	out := make(map[string]string)
	for i, val := range result {
		out[m.names[i]] = val
	}
	return out
}

// Expr returns the underlying regexp used for matching strings.
func (m *Template) Expr() string {
	if m == nil {
		return ""
	}
	return m.re.String()
}

// Names returns the list of named capture groups.
// The first item in the list will always be the empty string, which represents
// the implicit full match.
func (m *Template) Names() []string {
	if m == nil {
		return nil
	}
	return m.names
}
