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

var imgWidth = 512
var imgHeight = 512

var img = image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

var bg = color.RGBA{0, 0, 0, 255}

var shininess = 100.00

func main() {
	spheres := []raktracer.Sphere{
		raktracer.NewSphere(raktracer.Vector{75, 75, 450}, 100),
		raktracer.NewSphere(raktracer.Vector{-75, 0, 550}, 100),
	}
	light := raktracer.Vector{256, 512, -500}
	//light := raktracer.Vector{0, 0, -1000}

	for x := 0; x < imgWidth; x++ {
		for y := 0; y < imgHeight; y++ {
			img.Set(x, y, bg)
		}
	}

	camPos := raktracer.Vector{0, 0, -10000}

	for x := -imgWidth / 2; x < imgWidth/2; x++ {
		for y := -imgHeight / 2; y < imgHeight/2; y++ {
			screenPos := raktracer.Vector{float64(x), float64(y), 0}
			camVector := screenPos.Subtract(camPos).Normalize()

			r := raktracer.Ray{camPos, camVector}
			fmt.Printf("r: %s\n", r)

			var hitDist = -1.00
			var hitSphere raktracer.Sphere
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
			intersect := camPos.Add(r.Dir.Scale(hitDist))

			lightVector := light.Subtract(intersect).Normalize()

			n := intersect.Subtract(hitSphere.Pos).Normalize()

			lightIntersection := false
			lightRay := raktracer.Ray{intersect.Add(n.Scale(0.0001)), lightVector}
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
