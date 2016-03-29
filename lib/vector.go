package raktracer

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func (a Vector) String() string {
	return VectorString(a)
}

func (a Vector) Length() float64 {
	return VectorLength(a)
}

func (a Vector) Add(b Vector) Vector {
	return VectorsAdd(a, b)
}

func (a Vector) Subtract(b Vector) Vector {
	return VectorsSubtract(a, b)
}

func (a Vector) Scale(r float64) Vector {
	return VectorScaleByScalar(a, r)
}

func (a Vector) Normalize() Vector {
	return VectorNormalize(a)
}

func (a Vector) Dot(b Vector) float64 {
	return VectorsDotProduct(a, b)
}

func (a Vector) Cross(b Vector) Vector {
	return VectorsCrossProduct(a, b)
}

func VectorString(a Vector) string {
	return fmt.Sprintf("[%.2f %.2f %.2f]", a.X, a.Y, a.Z)
}

func VectorLength(a Vector) float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func VectorsAdd(a Vector, b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func VectorsSubtract(a Vector, b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func VectorScaleByScalar(a Vector, r float64) Vector {
	return Vector{a.X * r, a.Y * r, a.Z * r}
}

func VectorNormalize(a Vector) Vector {
	factor := 1.0 / a.Length()
	return Vector{a.X * factor, a.Y * factor, a.Z * factor}
}

func VectorsDotProduct(a Vector, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func VectorsCrossProduct(a Vector, b Vector) Vector {
	return Vector{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}
