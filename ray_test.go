package raktracer

import (
	"testing"
)

func TestRayString(t *testing.T) {
	cases := []struct {
		r    ray
		want string
	}{
		{
			ray{vector{0, 0, 0}, vector{0, 0, 0}},
			"ray{pos:[0.00 0.00 0.00] dir:[0.00 0.00 0.00]}",
		},
		{
			ray{vector{1, 0.5, 0.99}, vector{1, 2, 3}},
			"ray{pos:[1.00 0.50 0.99] dir:[1.00 2.00 3.00]}",
		},
	}
	for _, c := range cases {
		got := c.r.String()
		if got != c.want {
			t.Errorf("%s.String() expected %s, got %s", c.r, c.want, got)
		}
	}
}
