package raktracer

import (
	"fmt"
	"math"
)

// A Vector represents a vector in euclidian space (right hand orientation).
type Vector struct {
	X, Y, Z float64
}

// String returns a string representation of the vector a.
func (a Vector) String() string {
	return fmt.Sprintf("[%.2f %.2f %.2f]", a.X, a.Y, a.Z)
}

// Length returns the lenght of the euclidian vector a.
func (a Vector) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Add returns  a new vector which is the sum of the euclidian vectors a and b.
func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Subtract returns a new vector which is the difference between the euclidian
// vectors a and b.
func (a Vector) Subtract(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Scale returns a new vector where each component of the vector a has been
// multiplied by the scalar r.
func (a Vector) Scale(r float64) Vector {
	return Vector{a.X * r, a.Y * r, a.Z * r}
}

// Normalize returns a new unit vector, vector with length 1,  with the same
// directionaly as a.
func (a Vector) Normalize() Vector {
	factor := 1.0 / a.Length()
	return Vector{a.X * factor, a.Y * factor, a.Z * factor}
}

// Dot returns a new vector which is the dot product of a and b.
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross returns a new vector which is the cross product of a and b.
func (a Vector) Cross(b Vector) Vector {
	return Vector{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}
