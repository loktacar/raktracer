package raktracer

import (
	"fmt"
	"math"
)

type vector struct {
	x, y, z float64
}

func (a vector) Length() float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)
}

func (a vector) String() string {
	return fmt.Sprintf("[%.2f %.2f %.2f]", a.x, a.y, a.z)
}

/***** Functions with work with vectors *****/
func Add(a vector, b vector) vector {
	return vector{a.x + b.x, a.y + b.y, a.z + b.z}
}

func Scale(a vector, r float64) vector {
	return vector{a.x * r, a.y * r, a.z * r}
}

func Normalize(a vector) vector {
	factor := 1.0 / a.Length()
	return vector{a.x * factor, a.y * factor, a.z * factor}
}

func DotProduct(a vector, b vector) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func CrossProduct(a vector, b vector) vector {
	return vector{a.y*b.z - a.z*b.y, a.z*b.x - a.x*b.z, a.x*b.y - a.y*b.x}
}
