package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	cases := []struct { //declare an array of struct
		in, want string //with two properties of type string
	}{ //initialise array by adding cases
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases { //foreach c in cases
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
