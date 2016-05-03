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
			"Ray{Pos:[0.00 0.00 0.00] Dir:[0.00 0.00 0.00]}",
		},
		{
			Ray{Vector{1, 0.5, 0.99}, Vector{1, 2, 3}},
			"Ray{Pos:[1.00 0.50 0.99] Dir:[1.00 2.00 3.00]}",
		},
	}
	for _, c := range cases {
		got := RayString(c.r)
		if got != c.want {
			t.Errorf("%s.String() = %s, want %s", c.r, got, c.want)
		}
	}
	for _, c := range cases {
		got := c.r.String()
		if got != c.want {
			t.Errorf("%s.String() = %s, want %s", c.r, got, c.want)
		}
	}
}
