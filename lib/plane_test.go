package raktracer

import (
	"testing"
)

func TestPlaneConstructor(t *testing.T) {
	cases := []struct {
		p    Vector
		n    Vector
		want Plane
	}{
		{Vector{0, 0, 0}, Vector{0, 0, 1}, Plane{Vector{0, 0, 0}, Vector{0, 0, 1}}},
	}
	for _, c := range cases {
		got := NewPlane(c.p, c.n)
		if got != c.want {
			t.Errorf("PlaneNew(%s, %s) = %s, want %s", c.p, c.n, got, c.want)
		}
	}
}

func TestPlaneString(t *testing.T) {
	cases := []struct {
		p    Plane
		want string
	}{
		{Plane{Vector{0, 0, 0}, Vector{0, 0, 1}}, "Plane{P:[0.00 0.00 0.00] N:[0.00 0.00 1.00]}"},
	}
	for _, c := range cases {
		got := PlaneString(c.p)
		if got != c.want {
			t.Errorf("PlaneString(%s) = %s, want %s", c.p, got, c.want)
		}
	}
	for _, c := range cases {
		got := c.p.String()
		if got != c.want {
			t.Errorf("%s.String() = %s, want %s", c.p, got, c.want)
		}
	}
}

func TestPlaneIntersects(t *testing.T) {
	cases := []struct {
		p            Plane
		r            Ray
		wantHit      bool
		wantDistance float64
	}{
		{Plane{Vector{0, 0, 0}, Vector{0, 0, 1}}, Ray{Vector{0, 0, 0}, Vector{0, 0, 1}}, true, 0.00},
		{Plane{Vector{0, 0, 0}, Vector{0, 0, 1}}, Ray{Vector{1, 0, 0}, Vector{1, 0, 0}}, false, 0.00},
		{Plane{Vector{0, 0, 0}, Vector{0, 0, 1}}, Ray{Vector{0, 0, -1}, Vector{0, 0, 1}}, true, 1.00},
		{Plane{Vector{0, 0, 0}, Vector{0, 0, -1}}, Ray{Vector{0, 0, 1}, Vector{0, 0, -1}}, true, 1.00},
		{Plane{Vector{0, 0, 0}, Vector{0, 0, -1}}, Ray{Vector{0, 0, 1}, Vector{0, 0, 1}}, false, 0.00},
	}
	for _, c := range cases {
		gotHit, gotDistance := c.p.Intersects(c.r)
		if gotHit != c.wantHit || gotDistance != c.wantDistance {
			t.Errorf("%s.Intersects(%s) = %t, %.2f; want %t, %.2f", c.p, c.r, gotHit, gotDistance, c.wantHit, c.wantDistance)
		}
	}
}
