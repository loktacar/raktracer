package raktracer

import (
	"fmt"
	"math"
)

type vector struct {
	x, y, z float64
}

func (a vector) String() string {
	return VectorString(a)
}

func (a vector) Length() float64 {
	return VectorLength(a)
}

func (a vector) Add(b vector) vector {
	return VectorsAdd(a, b)
}

func (a vector) Subtract(b vector) vector {
	return VectorsSubtract(a, b)
}

func (a vector) Scale(r float64) vector {
	return VectorScaleByScalar(a, r)
}

func (a vector) Normalize() vector {
	return VectorNormalize(a)
}

func (a vector) Dot(b vector) float64 {
	return VectorsDotProduct(a, b)
}

func (a vector) Cross(b vector) vector {
	return VectorsCrossProduct(a, b)
}

func VectorString(a vector) string {
	return fmt.Sprintf("[%.2f %.2f %.2f]", a.x, a.y, a.z)
}

func VectorLength(a vector) float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)
}

func VectorsAdd(a vector, b vector) vector {
	return vector{a.x + b.x, a.y + b.y, a.z + b.z}
}

func VectorsSubtract(a vector, b vector) vector {
	return vector{a.x - b.x, a.y - b.y, a.z - b.z}
}

func VectorScaleByScalar(a vector, r float64) vector {
	return vector{a.x * r, a.y * r, a.z * r}
}

func VectorNormalize(a vector) vector {
	factor := 1.0 / a.Length()
	return vector{a.x * factor, a.y * factor, a.z * factor}
}

func VectorsDotProduct(a vector, b vector) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func VectorsCrossProduct(a vector, b vector) vector {
	return vector{a.y*b.z - a.z*b.y, a.z*b.x - a.x*b.z, a.x*b.y - a.y*b.x}
}
