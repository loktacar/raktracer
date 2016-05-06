package raktracer

import (
	"fmt"
)

// A Plane represents a plane in euclidian space, defined by a point on the
// plane and the normal of the plane (giving the orientation of the plane).
type Plane struct {
	Pos  Vector
	Norm Vector
}

func (p Plane) String() string {
	return PlaneString(p)
}

func NewPlane(p, n Vector) Plane {
	return Plane{p, n.Normalize()}
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

func (p Plane) NormalVector(pos Vector) Vector {
	return p.Norm.Scale(-1)
}

func PlaneString(p Plane) string {
	return fmt.Sprintf("Plane{Pos:%s Norm:%s}", p.Pos, p.Norm)
}
