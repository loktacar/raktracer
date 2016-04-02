package raktracer

import (
	"testing"
)

func TestSphereString(t *testing.T) {
	cases := []struct {
		s    Sphere
		want string
	}{
		{Sphere{Vector{0, 0, 0}, 0}, "Sphere{Pos:[0.00 0.00 0.00] R:0.00}"},
		{Sphere{Vector{1, 2, 3}, 0.99}, "Sphere{Pos:[1.00 2.00 3.00] R:0.99}"},
	}
	for _, c := range cases {
		got := SphereString(c.s)
		if got != c.want {
			t.Errorf("SphereString(%s) expected %s, got %s", c.s, c.want, got)
		}
	}
	for _, c := range cases {
		got := c.s.String()
		if got != c.want {
			t.Errorf("%s.String() expected %s, got %s", c.s, c.want, got)
		}
	}
}

func TestSphereIntersects(t *testing.T) {
	cases := []struct {
		s            Sphere
		r            Ray
		wantHit      bool
		wantDistance float64
	}{
		{Sphere{Vector{0, 0, 10}, 1}, Ray{Vector{0, 0, 0}, Vector{0, 0, 1}}, true, 9},
		{Sphere{Vector{0, 0, 10}, 1}, Ray{Vector{0, 0, 20}, Vector{0, 0, -1}}, true, 9},
		{Sphere{Vector{0, 0, 10}, 1}, Ray{Vector{0, 0, 0}, Vector{0, 1, 0}}, false, 0},
		{Sphere{Vector{0, 0, 0}, 4}, Ray{Vector{0, 5, 0}, Vector{0, -1, 0}}, true, 1},
		{Sphere{Vector{0, 0, 0}, 20}, Ray{Vector{0, 5, 0}, Vector{0, -1, 0}}, false, 0},
	}
	for _, c := range cases {
		gotHit, gotDistance := c.s.Intersects(c.r)
		if gotHit != c.wantHit || gotDistance != c.wantDistance {
			t.Errorf("%s.intersects(%s) expected %t, %.2f; got %t, %.2f", c.s, c.r, c.wantHit, c.wantDistance, gotHit, gotDistance)
		}
	}
}
