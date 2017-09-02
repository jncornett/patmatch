package patmatch

import "testing"

func TestInterpolate(t *testing.T) {
	tests := []struct {
		s      string
		vals   map[string]string
		expect string
	}{
		{
			s:      "Hello, [name]",
			vals:   map[string]string{"name": "world"},
			expect: "Hello, world",
		},
		{
			s:      "Hello, [name]",
			vals:   nil,
			expect: "Hello, ",
		},
		{
			s:      "Hello",
			vals:   map[string]string{"name": "world"},
			expect: "Hello",
		},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			actual := Interpolate(test.s, test.vals)
			if test.expect != actual {
				t.Errorf("want %q, got %q", test.expect, actual)
			}
		})
	}
}

func TestInterpolatePartial(t *testing.T) {
	tests := []struct {
		s      string
		vals   map[string]string
		expect string
	}{
		{
			s:      "Hello, [name]",
			vals:   map[string]string{"name": "world"},
			expect: "Hello, world",
		},
		{
			s:      "Hello, [name]",
			vals:   nil,
			expect: "Hello, [name]",
		},
		{
			s:      "Hello",
			vals:   map[string]string{"name": "world"},
			expect: "Hello",
		},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			actual := InterpolatePartial(test.s, test.vals)
			if test.expect != actual {
				t.Errorf("want %q, got %q", test.expect, actual)
			}
		})
	}
}
