package raktracer

import (
	"fmt"
)

type Sphere struct {
	Pos Vector
	R   float64
}

func (s Sphere) String() string {
	return SphereString(s)
}

func (s Sphere) Intersects(r Ray) (bool, float64) {
	// Algorithm source: http://www.scratchapixel.com/lessons/3d-basic-rendering/minimal-ray-tracer-rendering-simple-shapes/minimal-ray-tracer-rendering-spheres
	l := r.Pos.Subtract(s.Pos)

	a := r.Dir.Dot(r.Dir)
	b := 2 * r.Dir.Dot(l)
	c := l.Dot(l) - s.R*s.R

	// fmt.Printf("l: %.2f, a: %.2f, b: %.2f, c: %.2f\n", l, a, b, c)

	solved, t0, t1 := SolveQuadratic(a, b, c)

	// fmt.Printf("solved: %t, t0: %.2f, t1: %.2f\n", solved, t0, t1)

	if !solved {
		// fmt.Printf("Wasn't solved")
		return false, 0
	}

	if t0 < t1 {
		k := t0
		t0 = t1
		t1 = k
		if t0 < 0 {
			// fmt.Printf("Something? %.2f, %.2f\n", t0, t1)
			return false, 0
		}
	}

	return true, t0
}

func SphereString(s Sphere) string {
	return fmt.Sprintf("Sphere{Pos:%s R:%.2f}", s.Pos, s.R)
}
