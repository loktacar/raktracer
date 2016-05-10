package raktracer

type Object interface {
	// Intersects returns true if the ray r intersects this object, if the ray
	// intersects then dist is the distance from the origin of the ray r to the
	// closest intersection point.
	Intersects(r Ray) (intersects bool, dist float64)

	// SurfaceProperties returns the surface normal (norm) of the object at the
	// point p, the reflection direction with respect to the view direction
	// (vDir). Diffuse, Specular and Reflection coefficients (dC, sC, and rC
	// respectively) and the specular intensity (sN).
	SurfaceProperties(p Vector, vDir Vector) (norm Vector, refDir Vector, dC float64, sC float64, sN float64, rC float64)
}
