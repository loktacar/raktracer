package raktracer

import (
	"testing"
)

func TestSphereConstructor(t *testing.T) {
	cases := []struct {
		v    Vector
		r    float64
		dC   float64
		sC   float64
		sN   float64
		rC   float64
		want Sphere
	}{
		{Vector{0, 0, 0}, 0, 1, 1, 1, 1, Sphere{Vector{0, 0, 0}, 0, 1, 1, 1, 1, 0}},
		{Vector{1, 1, 1}, 1, 1, 1, 1, 1, Sphere{Vector{1, 1, 1}, 1, 1, 1, 1, 1, 1}},
		{Vector{55, 55, 55}, 20, 1, 1, 1, 1, Sphere{Vector{55, 55, 55}, 20, 1, 1, 1, 1, 400}},
	}
	for _, c := range cases {
		got := NewSphere(c.v, c.r, c.dC, c.sC, c.sN, c.rC)
		if got != c.want {
			t.Errorf("SphereNew(%s, %.2f) = %s, want %s", c.v, c.r, got, c.want)
		}
	}
}

func TestSphereString(t *testing.T) {
	cases := []struct {
		s    Sphere
		want string
	}{
		{Sphere{Vector{0, 0, 0}, 0, 1, 1, 1, 1, 0}, "Sphere{Pos:[0.00 0.00 0.00] R:0.00}"},
		{Sphere{Vector{1, 2, 3}, 0.99, 1, 1, 1, 1, 0.9801}, "Sphere{Pos:[1.00 2.00 3.00] R:0.99}"},
	}
	for _, c := range cases {
		got := c.s.String()
		if got != c.want {
			t.Errorf("%s.String() = %s, want %s", c.s, got, c.want)
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
		{Sphere{Vector{0, 0, 10}, 1, 1, 1, 1, 1, 1}, Ray{Vector{0, 0, 0}, Vector{0, 0, 1}}, true, 9},
		{Sphere{Vector{0, 0, 10}, 1, 1, 1, 1, 1, 1}, Ray{Vector{0, 0, 20}, Vector{0, 0, -1}}, true, 9},
		{Sphere{Vector{0, 0, 10}, 1, 1, 1, 1, 1, 1}, Ray{Vector{0, 0, 0}, Vector{0, 1, 0}}, false, 0},
		{Sphere{Vector{0, 0, 0}, 4, 1, 1, 1, 1, 16}, Ray{Vector{0, 5, 0}, Vector{0, -1, 0}}, true, 1},
		{Sphere{Vector{0, 0, 0}, 20, 1, 1, 1, 1, 400}, Ray{Vector{0, 5, 0}, Vector{0, -1, 0}}, false, 0},
	}
	for _, c := range cases {
		gotHit, gotDistance := c.s.Intersects(c.r)
		if gotHit != c.wantHit || gotDistance != c.wantDistance {
			t.Errorf("%s.Intersects(%s) = %t, %.2f; want %t, %.2f", c.s, c.r, gotHit, gotDistance, c.wantHit, c.wantDistance)
		}
	}
}
