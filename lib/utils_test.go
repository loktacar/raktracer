package raktracer

import (
	"testing"
)

func TestSolveQuadratic(t *testing.T) {
	cases := []struct {
		a, b, c float64
		solved  bool
		x0, x1  float64
	}{
		{1, 3, -4, true, -4, 1},                    // Two solutions
		{9, 12, 4, true, -(2.00 / 3), -(2.00 / 3)}, // One solution
		{3, 4, 2, false, 0.00, 0.00},               // No solutions
	}
	for _, c := range cases {
		gotSolved, got1, got2 := SolveQuadratic(c.a, c.b, c.c)
		if gotSolved != c.solved || got1 != c.x0 || got2 != c.x1 {
			t.Errorf("SolveQuadratic(%f, %f, %f) expected %t, %f, %f; got %t, %f, %f.", c.a, c.b, c.c, c.solved, c.x0, c.x1, gotSolved, got1, got2)
		}
	}
}
