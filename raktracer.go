package main

import (
	//"fmt"
	"github.com/loktacar/raktracer/lib"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var img = image.NewRGBA(image.Rect(0, 0, 512, 512))

var black = color.RGBA{0, 0, 0, 255}

var diffuseCoefficient = 0.95

func main() {
	spheres := []raktracer.Sphere{
		raktracer.Sphere{raktracer.Vector{75, 0, 550}, 100},
		raktracer.Sphere{raktracer.Vector{-75, 0, 450}, 100},
	}
	light := raktracer.Vector{256, 512, -500}
	//light := raktracer.Vector{0, 0, -1000}

	for x := 0; x < 512; x++ {
		for y := 0; y < 512; y++ {
			img.Set(x, y, black)
		}
	}

	for x := -256; x < 256; x++ {
		for y := -256; y < 256; y++ {
			camPos := raktracer.Vector{float64(x), float64(y), 0}

			r := raktracer.Ray{camPos, raktracer.Vector{0, 0, 1}}
			for _, s := range spheres {
				i, dist := s.Intersects(r)
				if i {
					intersect := camPos.Add(r.Dir.Scale(dist))

					lV := light.Subtract(intersect).Normalize()

					lightIntersection := false
					lR := raktracer.Ray{intersect, lV}
					for _, s2 := range spheres {
						if s2 != s {
							iL, _ := s2.Intersects(lR)
							lightIntersection = iL || lightIntersection
						}
					}

					diffLightValue := 0.05
					if !lightIntersection {
						n := s.Pos.Subtract(intersect).Normalize()

						diffLightValue += diffuseCoefficient * math.Max(0, n.Dot(lV)) * 0.95
					}
					//fmt.Printf("%s . %s = %.2f\n", lV, n, diffLightValue)
					c := color.RGBA{
						uint8(255 * diffLightValue),
						uint8(255 * diffLightValue),
						uint8(255 * diffLightValue),
						255}

					img.Set(x+255, y+255, c)
				}
			}
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
