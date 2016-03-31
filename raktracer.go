package main

import (
	"fmt"
	"github.com/loktacar/raktracer/lib"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var imgWidth = 1024
var imgHeight = 1024

var img = image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

var black = color.RGBA{0, 0, 0, 255}

var diffuseCoefficient = 0.95

func main() {
	spheres := []raktracer.Sphere{
		raktracer.Sphere{raktracer.Vector{75, 0, 550}, 100},
		raktracer.Sphere{raktracer.Vector{-75, 0, 450}, 100},
	}
	light := raktracer.Vector{256, 512, -500}
	//light := raktracer.Vector{0, 0, -1000}

	for x := 0; x < imgWidth; x++ {
		for y := 0; y < imgHeight; y++ {
			img.Set(x, y, black)
		}
	}

	for x := -imgWidth / 2; x < imgWidth/2; x++ {
		for y := -imgHeight / 2; y < imgHeight/2; y++ {
			camPos := raktracer.Vector{float64(x), float64(y), 0}

			r := raktracer.Ray{camPos, raktracer.Vector{0, 0, 1}}

			var hitDist = -1.00
			var hitSphere raktracer.Sphere
			var intersect raktracer.Vector

			for _, s := range spheres {
				i, dist := s.Intersects(r)
				if i && (hitDist == -1.00 || dist < hitDist) {
					hitDist = dist
					hitSphere = s
					intersect = camPos.Add(r.Dir.Scale(dist))
				}
			}

			lightVector := light.Subtract(intersect).Normalize()

			lightIntersection := false
			lightRay := raktracer.Ray{intersect, lightVector}
			for _, s2 := range spheres {
				if s2 != hitSphere {
					iL, _ := s2.Intersects(lightRay)
					lightIntersection = iL || lightIntersection
				}
			}

			diffLightValue := 0.05
			if !lightIntersection {
				n := hitSphere.Pos.Subtract(intersect).Normalize()

				diffLightValue += diffuseCoefficient * math.Max(0, n.Dot(lightVector)) * 0.95
			}
			c := color.RGBA{
				uint8(255 * diffLightValue),
				uint8(255 * diffLightValue),
				uint8(255 * diffLightValue),
				255}

			img.Set(x+imgWidth/2, y+imgHeight/2, c)
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)

	fmt.Println("Done!")
}
