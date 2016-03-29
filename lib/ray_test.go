package raktracer

import (
	"testing"
)

func TestRayString(t *testing.T) {
	cases := []struct {
		r    Ray
		want string
	}{
		{
			Ray{Vector{0, 0, 0}, Vector{0, 0, 0}},
			"Ray{pos:[0.00 0.00 0.00] dir:[0.00 0.00 0.00]}",
		},
		{
			Ray{Vector{1, 0.5, 0.99}, Vector{1, 2, 3}},
			"Ray{pos:[1.00 0.50 0.99] dir:[1.00 2.00 3.00]}",
		},
	}
	for _, c := range cases {
		got := c.r.String()
		if got != c.want {
			t.Errorf("%s.String() expected %s, got %s", c.r, c.want, got)
		}
	}
}
