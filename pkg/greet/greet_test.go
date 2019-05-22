package greet

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		res  string
	}{
		{"", "Hello "},
		{"name", "Hello name"},
	}

	for ind, test := range tests {
		res := Greet(test.name)
		if res != test.res {
			t.Errorf("test %d: expected '%s' received '%s'", ind, test.res, res)
		}
	}
}
