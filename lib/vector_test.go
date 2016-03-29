package raktracer

import (
	"math"
	"testing"
)

func TestVectorString(t *testing.T) {
	cases := []struct {
		v    Vector
		want string
	}{
		{Vector{0, 0, 0}, "[0.00 0.00 0.00]"},
		{Vector{1, 0.5, 0.99}, "[1.00 0.50 0.99]"},
	}
	for _, c := range cases {
		got := VectorString(c.v)
		if got != c.want {
			t.Errorf("VectorString(%s) expected %s, got %s", c.v, c.want, got)
		}
	}
	for _, c := range cases {
		got := c.v.String()
		if got != c.want {
			t.Errorf("%s.String() expected %s, got %s", c.v, c.want, got)
		}
	}
}

func TestVectorLength(t *testing.T) {
	cases := []struct {
		v    Vector
		want float64
	}{
		{Vector{0, 0, 0}, 0},
		{Vector{1, 0, 0}, 1},
		{Vector{0, 1, 0}, 1},
		{Vector{0, 0, 1}, 1},
		{Vector{5, 10, 10}, 15},
	}
	for _, c := range cases {
		got := VectorLength(c.v)
		if got != c.want {
			t.Errorf("VectorLength(%s) expected %f, got %f", c.v, c.want, got)
		}
	}
	for _, c := range cases {
		got := c.v.Length()
		if got != c.want {
			t.Errorf("%s.Length() expected %f, got %f", c.v, c.want, got)
		}
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b, want Vector
	}{
		{Vector{1, 2, 3}, Vector{-1, -2, -3}, Vector{0, 0, 0}},
		{Vector{-1, -2, -3}, Vector{1, 2, 3}, Vector{0, 0, 0}},
		{Vector{0, 0, 0}, Vector{0, 0, 0}, Vector{0, 0, 0}},
		{Vector{5, 5, 5}, Vector{10, 10, 10}, Vector{15, 15, 15}},
	}
	for _, c := range cases {
		got := VectorsAdd(c.a, c.b)
		if got != c.want {
			t.Errorf("VectorsAdd(%s, %s) expected %s, got %s", c.a, c.b, c.want, got)
		}
	}
	for _, c := range cases {
		got := c.a.Add(c.b)
		if got != c.want {
			t.Errorf("%s.Add(%s) expected %s, got %s", c.a, c.b, c.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		a, b, want Vector
	}{
		{Vector{1, 2, 3}, Vector{-1, -2, -3}, Vector{2, 4, 6}},
		{Vector{-1, -2, -3}, Vector{1, 2, 3}, Vector{-2, -4, -6}},
		{Vector{0, 0, 0}, Vector{0, 0, 0}, Vector{0, 0, 0}},
		{Vector{5, 5, 5}, Vector{10, 10, 10}, Vector{-5, -5, -5}},
	}
	for _, c := range cases {
		got := VectorsSubtract(c.a, c.b)
		if got != c.want {
			t.Errorf("VectorsSubtract(%s, %s) expected %s, got %s", c.a, c.b, c.want, got)
		}
	}
	for _, c := range cases {
		got := c.a.Subtract(c.b)
		if got != c.want {
			t.Errorf("%s.Subtract(%s) expected %s, got %s", c.a, c.b, c.want, got)
		}
	}
}

func TestScale(t *testing.T) {
	cases := []struct {
		v    Vector
		s    float64
		want Vector
	}{
		{Vector{0, 0, 0}, 5, Vector{0, 0, 0}},
		{Vector{1, 1, 1}, 5, Vector{5, 5, 5}},
		{Vector{-1, 0.5, 12}, 2, Vector{-2, 1, 24}},
	}
	for _, c := range cases {
		got := VectorScaleByScalar(c.v, c.s)
		if got != c.want {
			t.Errorf("VectorScaleByScalar(%s, %f) expected %s, got %s", c.v, c.s, c.want, got)
		}
	}
	for _, c := range cases {
		got := c.v.Scale(c.s)
		if got != c.want {
			t.Errorf("%s.Scale(%f) expected %s, got %s", c.v, c.s, c.want, got)
		}
	}
}

func TestNormalize(t *testing.T) {
	cases := []struct {
		v    Vector
		want Vector
	}{
		{Vector{1, 0, 0}, Vector{1, 0, 0}},
		{Vector{0, 1, 0}, Vector{0, 1, 0}},
		{Vector{0, 0, 1}, Vector{0, 0, 1}},
		{Vector{5, 10, 10}, Vector{5. / 15, 10. / 15, 10. / 15}},
	}
	for _, c := range cases {
		got := VectorNormalize(c.v)
		if got != c.want {
			t.Errorf("VectorNormalize(%s) expected %s, got %s", c.v, c.want, got)
		}
	}
	for _, c := range cases {
		got := c.v.Normalize()
		if got != c.want {
			t.Errorf("Normalize(%s) expected %s, got %s", c.v, c.want, got)
		}
	}

	v1 := Vector{0, 0, 0}
	vN := VectorNormalize(v1)
	if !math.IsNaN(vN.x) && !math.IsNaN(vN.y) && !math.IsNaN(vN.z) {
		t.Errorf("VectorNormalize(%s) expected NaN results, got %s", v1, vN)
	}
	vN2 := v1.Normalize()
	if !math.IsNaN(vN2.x) && !math.IsNaN(vN2.y) && !math.IsNaN(vN2.z) {
		t.Errorf("%s.Normalize() expected NaN results, got %s", v1, vN2)
	}
}

func TestDotProduct(t *testing.T) {
	cases := []struct {
		a    Vector
		b    Vector
		want float64
	}{
		{Vector{0, 0, 0}, Vector{0, 0, 0}, 0},
		{Vector{1, 2, 3}, Vector{1, 2, 3}, 14},
		{Vector{3, 3, 3}, Vector{3, 3, 3}, 27},
		{Vector{1, 0, -1}, Vector{1, 0, -1}, 2},
	}
	for _, c := range cases {
		got := VectorsDotProduct(c.a, c.b)
		if got != c.want {
			t.Errorf("VectorsDotProduct(%s, %s) expected %s, got %s", c.a, c.b, c.want, got)
		}
	}
	for _, c := range cases {
		got := c.a.Dot(c.b)
		if got != c.want {
			t.Errorf("%s.Dot(%s) expected %s, got %s", c.a, c.b, c.want, got)
		}
	}
}

func TestCrossProduct(t *testing.T) {
	cases := []struct {
		a    Vector
		b    Vector
		want Vector
	}{
		{Vector{0, 0, 0}, Vector{0, 0, 0}, Vector{0, 0, 0}},
		{Vector{1, 2, 3}, Vector{1, 2, 3}, Vector{0, 0, 0}},
		{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{-3, 6, -3}},
	}
	for _, c := range cases {
		got := VectorsCrossProduct(c.a, c.b)
		if got != c.want {
			t.Errorf("VectorsCrossProduct(%s, %s) expected %s, got %s", c.a, c.b, c.want, got)
		}
	}
	for _, c := range cases {
		got := c.a.Cross(c.b)
		if got != c.want {
			t.Errorf("%s.Cross(%s) expected %s, got %s", c.a, c.b, c.want, got)
		}
	}
}
