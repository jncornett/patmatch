package patmatch

import "testing"

func TestTemplate_Apply(t *testing.T) {
	templ, err := ParseFlags("%(foo)s and %(bar)s", 0)
	if err != nil {
		t.Fatal(err)
	}
	matches := templ.Apply("buzz and bat")
	if len(matches) != 3 {
		t.Fatalf("matches: want %d entries, got %d", 3, len(matches))
	}
	if matches[""] != "buzz and bat" {
		t.Errorf("full text match: want %q, got %q", "buzz and bat", matches[""])
	}
	if matches["foo"] != "buzz" {
		t.Errorf("foo match: want %q, got %q", "buzz", matches["foo"])
	}
	if matches["bar"] != "bat" {
		t.Errorf("bar match: want %q, got %q", "bat", matches["bar"])
	}
}
