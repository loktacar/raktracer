package raktracer

import (
	"fmt"
)

// A ray represents a ray in euclidian space, defined by a point vector and a
// direction vector.
type Ray struct {
	Pos, Dir Vector
}

// String returns a string representation of the ray r.
func (r Ray) String() string {
	return RayString(r)
}

// RayString returns a string representation of the ray r.
func RayString(r Ray) string {
	return fmt.Sprintf("Ray{Pos:%s Dir:%s}", r.Pos, r.Dir)
}
