package raktracer

import (
	"testing"
)

func TestSphereString(t *testing.T) {
	cases := []struct {
		s    sphere
		want string
	}{
		{sphere{vector{0, 0, 0}, 0}, "sphere{pos:[0.00 0.00 0.00] r:0.00}"},
		{sphere{vector{1, 2, 3}, 0.99}, "sphere{pos:[1.00 2.00 3.00] r:0.99}"},
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
		s            sphere
		r            ray
		wantHit      bool
		wantDistance float64
	}{
		{sphere{vector{0, 0, 10}, 1}, ray{vector{0, 0, 20}, vector{0, 0, -1}}, true, 9},
		{sphere{vector{0, 0, 10}, 1}, ray{vector{0, 0, 0}, vector{0, 0, 1}}, true, 9},
		{sphere{vector{0, 0, 10}, 1}, ray{vector{0, 0, 0}, vector{0, 1, 0}}, false, 0},
		{sphere{vector{0, 0, 0}, 4}, ray{vector{0, 5, 0}, vector{0, -1, 0}}, true, 1},
		{sphere{vector{0, 0, 0}, 20}, ray{vector{0, 5, 0}, vector{0, -1, 0}}, true, 0},
	}
	for _, c := range cases {
		gotHit, gotDistance := c.s.Intersects(c.r)
		if gotHit != c.wantHit && gotDistance != c.wantDistance {
			t.Errorf("%s.intersects(%s) expected %t, %.2f; got %t, %.2f", c.s, c.r, c.wantHit, c.wantDistance, gotHit, gotDistance)
		}
	}
}
