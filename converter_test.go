package ambrosio_distance_converter

import (
	"regexp"
	"testing"
)

func TestHandler(t *testing.T) {
	var tests = []struct {
		s    []string
		want string
	}{
		{[]string{"1 mile to km", "1", "mile", "to", "km"}, "1.60934"},
		{[]string{"1 mile to kilometer", "1", "mile", "to", "kilometer"}, "1.60934"},
		{[]string{"1 mile to kilometers", "1", "mile", "to", "kilometer"}, "1.60934"},

		{[]string{"1 meter to kilometers", "1", "meter", "to", "kilometer"}, "0.00100"},
		{[]string{"1 meter to kilometers", "1", "meter", "to", "kilometer"}, "0.00100"},
	}

	for _, c := range tests {
		got, _ := Converter.Handler(c.s)
		if got != c.want {
			t.Errorf("Converter.Handler(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}

func TestPattern(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"1 mile to kilometer", true},
		{"1.0 mile 2 kilometers", true},
		{"27.2 mile to km", true},
		{"13.3 meters to miles", true},
		{"13.3 meters to miles", true},

		{"1 seconds to miles", false},
		{"1f kilometer to miles", false},
	}

	for _, c := range tests {
		got, _ := regexp.MatchString(Converter.Pattern, c.s)
		if got != c.want {
			t.Errorf("regexp.MatchString(%q, %q) == %t, want %t", Converter.Pattern, c.s, got, c.want)
		}
	}
}
