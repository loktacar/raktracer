package raktracer

import (
	"fmt"
)

// A Plane represents a plane in euclidian space, defined by a point on the
// plane and the normal of the plane (giving the orientation of the plane).
type Plane struct {
	P Vector
	N Vector
}

func (p Plane) String() string {
	return PlaneString(p)
}

func NewPlane(p, n Vector) Plane {
	return Plane{p, n.Normalize()}
}

func (p Plane) Intersects(r Ray) (intersects bool, dist float64) {
	d := p.N.Dot(r.Dir)

	if d > 0.0000001 {
		p0l0 := p.P.Subtract(r.Pos)
		t := p0l0.Dot(p.N) / d
		return t >= 0, t
	}

	return false, 0
}

func (p Plane) NormalVector(pos Vector) Vector {
	return p.N.Scale(-1)
}

func PlaneString(p Plane) string {
	return fmt.Sprintf("Plane{P:%s N:%s}", p.P, p.N)
}
