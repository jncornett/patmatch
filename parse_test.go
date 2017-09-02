package patmatch

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in    string
		re    string
		names string
	}{
		{
			in:    "",
			re:    "",
			names: "",
		},
		{
			in:    "fixed",
			re:    "fixed",
			names: "",
		},
		{
			in:    "%s",
			re:    `.+`,
			names: "",
		},
		{
			in:    "%%",
			re:    `%`,
			names: "",
		},
		{
			in:    "%(named)s",
			re:    `(.+)`,
			names: ",named",
		},
		{
			in:    "   ",
			re:    `\s+`,
			names: "",
		},
		{
			in:    "twas %(time)s, and",
			re:    `twas\s+(.+),\s+and`,
			names: ",time",
		},
		{
			in:    "%(foo)s and %(bar)s",
			re:    `(.+)\s+and\s+(.+)`,
			names: ",foo,bar",
		},
		{
			in:    "%s and %s",
			re:    `.+\s+and\s+.+`,
			names: "",
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			m, err := ParseFlags(test.in, 0)
			if err != nil {
				t.Fatal(err)
			}
			re := m.Expr()
			if test.re != re {
				t.Log(re)
				t.Errorf("want %q, got %q", test.re, re)
			}
			names := strings.Join(m.names, ",")
			if test.names != names {
				t.Errorf("want %q, got %q", test.names, names)
			}
		})
	}
}
