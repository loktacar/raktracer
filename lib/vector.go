package raktracer

import (
	"fmt"
	"math"
)

type Vector struct {
	x, y, z float64
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
	return fmt.Sprintf("[%.2f %.2f %.2f]", a.x, a.y, a.z)
}

func VectorLength(a Vector) float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)
}

func VectorsAdd(a Vector, b Vector) Vector {
	return Vector{a.x + b.x, a.y + b.y, a.z + b.z}
}

func VectorsSubtract(a Vector, b Vector) Vector {
	return Vector{a.x - b.x, a.y - b.y, a.z - b.z}
}

func VectorScaleByScalar(a Vector, r float64) Vector {
	return Vector{a.x * r, a.y * r, a.z * r}
}

func VectorNormalize(a Vector) Vector {
	factor := 1.0 / a.Length()
	return Vector{a.x * factor, a.y * factor, a.z * factor}
}

func VectorsDotProduct(a Vector, b Vector) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func VectorsCrossProduct(a Vector, b Vector) Vector {
	return Vector{a.y*b.z - a.z*b.y, a.z*b.x - a.x*b.z, a.x*b.y - a.y*b.x}
}
