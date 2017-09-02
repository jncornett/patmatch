package patmatch

import (
	"regexp"
	"strings"
)

type regexpBuilder struct {
	exact bool
	parts []string
}

func newRegexpBuilder(exact, ignoreCase bool) regexpBuilder {
	var parts []string
	if ignoreCase {
		parts = append(parts, `(?i)`)
	}
	if exact {
		parts = append(parts, `^`)
	}
	return regexpBuilder{exact: exact, parts: parts}
}

func (b *regexpBuilder) addExpr(expr string) {
	b.parts = append(b.parts, expr)
}

func (b *regexpBuilder) addCapture(expr string) {
	b.parts = append(b.parts, `(`+expr+`)`)
}

func (b *regexpBuilder) addExact(s string) {
	b.parts = append(b.parts, regexp.QuoteMeta(s))
}

func (b regexpBuilder) build() *regexp.Regexp {
	if b.exact {
		b.parts = append(b.parts, `$`)
	}
	return regexp.MustCompile(strings.Join(b.parts, ""))
}
