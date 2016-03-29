package raktracer

import (
	"math"
)

func SolveQuadratic(a float64, b float64, c float64) (bool, float64, float64) {
	// Algorithm source: http://www.scratchapixel.com/lessons/3d-basic-rendering/minimal-ray-tracer-rendering-simple-shapes/minimal-ray-tracer-rendering-spheres
	var x0, x1, discr float64

	discr = b*b - 4*a*c

	if discr < 0 {
		return false, x0, x1
	} else if discr == 0 {
		x0 = (-0.5 * b / a)
		x1 = x0
	} else {
		var q float64
		if b > 0 {
			q = -0.5 * (b + math.Sqrt(discr))
		} else {
			q = -0.5 * (b - math.Sqrt(discr))
		}
		x0 = q / a
		x1 = c / q
	}

	if x0 > x1 {
		return true, x1, x0
	} else {
		return true, x0, x1
	}
}
