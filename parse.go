package patmatch

import (
	"errors"
	"regexp"
)

// ParseFlag represents a boolean flag that can be passed int ParseFlag.
type ParseFlag uint

const (
	// FlagNone represents no special options.
	FlagNone ParseFlag = 0
	// FlagWholeString forces the template to match the entire string.
	// By default the template matches on substrings.
	FlagWholeString = 1 << 0
	// FlagIgnoreCase causes the template to ignore any casing in the expression.
	FlagIgnoreCase = 1 << 1
)

// Parse parses a template expression and compiles it into a Template.
func Parse(expr string) (*Template, error) {
	return ParseFlags(expr, FlagNone)
}

// ParseFlags parses a template expression and compiles it into a Template,
// passing any any optional flags.
func ParseFlags(expr string, flags ParseFlag) (*Template, error) {
	builder := newRegexpBuilder(flags&FlagWholeString != 0, flags&FlagIgnoreCase != 0)
	names := []string{""}
	for _, group := range patExpr.FindAllStringSubmatch(expr, -1) {
		if group[patName] != "" {
			// named group (gets captured)
			verb, exists := verbMap[group[patNamedVerb]]
			if !exists {
				return nil, errors.New("unrecognized verb in group " + group[patName] + ": " + group[patNamedVerb])
			}
			builder.addCapture(verb)
			names = append(names, group[patName])
		} else if group[patVerb] == "%" {
			// escape sequence for '%'
			builder.addExact("%")
		} else if group[patVerb] != "" {
			// anonymous group (no capture)
			verb, exists := verbMap[group[patVerb]]
			if !exists {
				return nil, errors.New("unrecognized verb: " + group[patVerb])
			}
			builder.addExpr(verb)
		} else if group[patString] != "" {
			// exact match
			builder.addExact(group[patString])
		} else if group[patWs] != "" {
			// collapse whitespace
			builder.addExpr(`\s+`)
		}
	}
	return &Template{
		re:    builder.build(),
		names: names,
	}, nil
}

const (
	patName      = 1
	patNamedVerb = 2
	patVerb      = 3
	patString    = 4
	patWs        = 5
)

var (
	patExpr = regexp.MustCompile(`%\((\w+)\)(\w)|%([%\w])|([^%\s]+?)|(\s+)`)
	// Currently, only '%s' is supported, but we may support more verbs in the future ('%d' for numbers, for example).
	verbMap = map[string]string{
		"s": `.+`,
	}
)
