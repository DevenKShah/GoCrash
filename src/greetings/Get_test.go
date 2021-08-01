package greetings

import "testing"

func TestGet(t *testing.T) {
	cases := []struct {
		in       string
		want     string
		hasError bool
	}{
		{"Deven", "Hi, Deven. Welcome!", false},
		{"Suman", "Hi, Suman. Welcome!", false},
		{"", "", true},
	}

	for _, c := range cases {
		got, err := Get(c.in)
		hasError := err != nil

		if hasError != c.hasError {
			t.Error("Unexpected outcome")
		}
		if got != c.want {
			t.Errorf("Get(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
