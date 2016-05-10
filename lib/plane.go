package raktracer

import (
	"fmt"
)

// A Plane represents a plane in euclidian space, defined by a point on the
// plane and the normal of the plane (giving the orientation of the plane).
type Plane struct {
	Pos                   Vector
	Norm                  Vector
	DiffuseCoefficient    float64
	SpecularCoefficient   float64
	SpecularN             float64
	ReflectiveCoefficient float64
}

func (p Plane) String() string {
	return PlaneString(p)
}

func NewPlane(p, n Vector, dC float64, sC float64, sN float64) Plane {
	return Plane{p, n.Normalize(), dC, sC, sN, 0}
}

func (p Plane) Intersects(r Ray) (intersects bool, dist float64) {
	d := r.Dir.Dot(p.Norm)

	if d > 0.0000001 {
		p0l0 := p.Pos.Subtract(r.Pos)
		t := p0l0.Dot(p.Norm) / d
		return t >= 0, t
	}

	return false, 0
}

func (p Plane) SurfaceProperties(pos Vector, vDir Vector) (norm Vector, refDir Vector, dC float64, sC float64, sN float64, rC float64) {
	if vDir.Dot(p.Norm) > 0 {
		norm = p.Norm.Scale(-1)
	} else {
		norm = p.Norm
	}

	refDir = norm.Scale(2 * norm.Dot(vDir)).Subtract(vDir)

	return norm, refDir, p.DiffuseCoefficient, p.SpecularCoefficient, p.SpecularN, p.ReflectiveCoefficient
}

func PlaneString(p Plane) string {
	return fmt.Sprintf("Plane{Pos:%s Norm:%s}", p.Pos, p.Norm)
}
