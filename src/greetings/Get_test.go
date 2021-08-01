package greetings

import "testing"

func TestGet(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Deven", "Hi, Deven. Welcome!"},
		{"Suman", "Hi, Suman. Welcome!"},
	}

	for _, c := range cases {
		got := Get(c.in)
		if got != c.want {
			t.Errorf("Get(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}