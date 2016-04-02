package raktracer

import (
	"fmt"
	"math"
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
	l := s.Pos.Subtract(r.Pos)
	tca := l.Dot(r.Dir)
	if tca < 0 {
		return false, 0
	}
	d2 := l.Dot(l) - tca*tca
	if d2 > s.R*s.R {
		return false, 0
	}
	thc := math.Sqrt(s.R*s.R - d2)
	t0 := tca - thc
	t1 := tca + thc

	if t0 > t1 {
		t0 = t1
	}

	if t0 < 0 {
		return false, 0
	}

	return true, t0
}

func SphereString(s Sphere) string {
	return fmt.Sprintf("Sphere{Pos:%s R:%.2f}", s.Pos, s.R)
}
