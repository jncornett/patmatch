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
			names: "0",
		},
		{
			in:    "fixed",
			re:    "fixed",
			names: "0",
		},
		{
			in:    "%s",
			re:    `.+`,
			names: "0",
		},
		{
			in:    "%%",
			re:    `%`,
			names: "0",
		},
		{
			in:    "%(named)s",
			re:    `(.+)`,
			names: "0,named",
		},
		{
			in:    "   ",
			re:    `\s+`,
			names: "0",
		},
		{
			in:    "twas %(time)s, and",
			re:    `twas\s+(.+),\s+and`,
			names: "0,time",
		},
		{
			in:    "%(foo)s and %(bar)s",
			re:    `(.+)\s+and\s+(.+)`,
			names: "0,foo,bar",
		},
		{
			in:    "%s and %s",
			re:    `.+\s+and\s+.+`,
			names: "0",
		},
		{
			in:    "%()s and %()s",
			re:    `(.+)\s+and\s+(.+)`,
			names: "0,1,2",
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
