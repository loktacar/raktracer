package raktracer

import (
	"fmt"
)

type sphere struct {
	pos vector
	r   float64
}

func (s sphere) String() string {
	return SphereString(s)
}

func (s sphere) Intersects(r ray) (bool, float64) {
	// Algorithm source: http://www.scratchapixel.com/lessons/3d-basic-rendering/minimal-ray-tracer-rendering-simple-shapes/minimal-ray-tracer-rendering-spheres
	l := r.pos.Subtract(s.pos)

	a := r.dir.Dot(r.dir)
	b := 2 * r.dir.Dot(l)
	c := l.Dot(l) - s.r*s.r

	// fmt.Printf("l: %.2f, a: %.2f, b: %.2f, c: %.2f\n", l, a, b, c)

	solved, t0, t1 := SolveQuadratic(a, b, c)

	// fmt.Printf("solved: %t, t0: %.2f, t1: %.2f\n", solved, t0, t1)

	if !solved {
		// fmt.Printf("Wasn't solved")
		return false, 0
	}

	if t0 < 0 {
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

func SphereString(s sphere) string {
	return fmt.Sprintf("sphere{pos:%s r:%.2f}", s.pos, s.r)
}
