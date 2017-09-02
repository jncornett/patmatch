package patmatch

import (
	"testing"
)

func TestRegexBuilderAddExpr(t *testing.T) {
	test := "someExpr"
	b := newRegexpBuilder(false, false)
	b.addExpr(test)
	re := b.build()
	if test != re.String() {
		t.Errorf("want %q, got %q", test, re.String())
	}
}

func TestRegexBuilderAddExact(t *testing.T) {
	tests := []struct {
		in     string
		expect string
	}{
		{
			in:     "something",
			expect: `something`,
		},
		{
			in:     "%",
			expect: `%`,
		},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			b := newRegexpBuilder(false, false)
			b.addExact(test.in)
			re := b.build()
			if test.expect != re.String() {
				t.Errorf("want %q, got %q", test.expect, re.String())
			}
		})
	}
}

func TestRegexBuilderOptions(t *testing.T) {
	tests := []struct {
		exact      bool
		ignoreCase bool
		expect     string
	}{
		{true, false, "^$"},
		{true, true, "(?i)^$"},
		{false, true, "(?i)"},
	}
	for _, test := range tests {
		t.Run(test.expect, func(t *testing.T) {
			re := newRegexpBuilder(test.exact, test.ignoreCase).build()
			if test.expect != re.String() {
				t.Errorf("want %q, got %q", test.expect, re.String())
			}
		})
	}
}
