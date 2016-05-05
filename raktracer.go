package main

import (
	"fmt"
	"image/color"
	"math"

	. "github.com/loktacar/raktracer/lib"
)

var imgWidth = 1024
var imgHeight = 1024

var shininess = 100.00

func main() {
	spheres := []Sphere{
		NewSphere(Vector{150, -75, -650}, 200),
		NewSphere(Vector{-150, 0, -750}, 200),
	}
	light := Vector{256, -512, 500}
	//light := Vector{0, 0, -1000}

	cam := NewCamera(
		Vector{0, 0, 0},
		70,
		imgWidth,
		imgHeight)

	for point := range cam.ImagePoints() {
		r := point.SceneRay

		var hitDist = -1.00
		var hitSphere Sphere
		for _, s := range spheres {
			i, dist := s.Intersects(r)
			if i && (hitDist == -1.00 || dist < hitDist) {
				hitDist = dist
				hitSphere = s
			}
		}
		if hitDist == -1.00 {
			continue
		}
		intersect := r.Pos.Add(r.Dir.Scale(hitDist))

		lightVector := light.Subtract(intersect).Normalize()

		n := intersect.Subtract(hitSphere.Pos).Normalize()

		lightIntersection := false
		lightRay := Ray{intersect.Add(n.Scale(0.0001)), lightVector}
		for _, s2 := range spheres {
			iL, _ := s2.Intersects(lightRay)
			if iL {
				lightIntersection = true
				break
			}
		}

		diffLightValue := 0.05
		if !lightIntersection {
			// Ambient light 10%
			diffLightValue = 0.10
			reflectiveVector := n.Scale(2 * lightVector.Scale(-1).Dot(n)).Subtract(lightVector.Scale(-1)).Normalize()
			// Diffuse light 60%
			diffLightValue += 0.60 * math.Max(0, lightVector.Dot(n))
			// Specular light 20%
			diffLightValue += 0.20 * math.Pow(math.Max(0, reflectiveVector.Dot(r.Dir)), shininess)
		}
		c := color.RGBA{
			uint8(255 * diffLightValue),
			uint8(255 * diffLightValue),
			uint8(255 * diffLightValue),
			255}

		cam.SetPoint(point.ImageX, point.ImageY, c)
	}

	cam.Save("image.png")

	fmt.Println("Done!")
}
