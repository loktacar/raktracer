package raktracer

import (
	"fmt"
	"math"
)

// A Sphere represents a sphere in euclidian space, with a center position
// vector and a radius.
type Sphere struct {
	Pos                   Vector
	R                     float64
	DiffuseCoefficient    float64
	SpecularCoefficient   float64
	SpecularN             float64
	ReflectiveCoefficient float64
	r2                    float64
}

// String returns a string representation of the sphere s.
func (s Sphere) String() string {
	return fmt.Sprintf("Sphere{Pos:%s R:%.2f}", s.Pos, s.R)
}

// NewSphere returns a new sphere with the given position pos and radius r.
func NewSphere(pos Vector, r float64, dC float64, sC float64, sN float64, rC float64) Sphere {
	return Sphere{pos, r, dC, sC, sN, rC, r * r}
}

// Intersects checks if the ray r intersects with the sphere s. If the ray
// intersects the returned intersects value is returned as true, otherwise it
// is false. If the ray intersects the distance from the ray origin to the
// first intersection point is returned as dist.
func (s Sphere) Intersects(r Ray) (intersects bool, dist float64) {
	// Algorithm source: http://www.scratchapixel.com/lessons/3d-basic-rendering/minimal-ray-tracer-rendering-simple-shapes/minimal-ray-tracer-rendering-spheres
	l := s.Pos.Subtract(r.Pos)
	tca := l.Dot(r.Dir)
	if tca < 0 {
		return false, 0
	}
	d2 := l.Dot(l) - tca*tca
	if d2 > s.r2 {
		return false, 0
	}
	thc := math.Sqrt(s.r2 - d2)
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

func (s Sphere) NormalVector(pos Vector) Vector {
	return pos.Subtract(s.Pos).Normalize()
}

func (s Sphere) SurfaceProperties(pos Vector, vDir Vector) (norm Vector, refDir Vector, dC float64, sC float64, sN float64, rC float64) {
	norm = s.NormalVector(pos)
	refDir = norm.Scale(2 * norm.Dot(vDir)).Subtract(vDir)

	return norm, refDir, s.DiffuseCoefficient, s.SpecularCoefficient, s.SpecularN, s.ReflectiveCoefficient
}
