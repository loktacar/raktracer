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

var img = image.NewRGBA(image.Rect(0, 0, 512, 512))

var black = color.RGBA{0, 0, 0, 255}

var diffuseCoefficient = 0.95

func main() {
	s := raktracer.Sphere{raktracer.Vector{0, 0, 50}, 100}
	light := raktracer.Vector{256, 512, -500}
	//light := raktracer.Vector{0, 0, -1000}

	for x := -256; x < 256; x++ {
		for y := -256; y < 256; y++ {
			camPos := raktracer.Vector{float64(x), float64(y), 0}

			r := raktracer.Ray{camPos, raktracer.Vector{0, 0, 1}}
			i, dist1, _ := s.Intersects(r)
			if i {
				intersect := camPos.Add(r.Dir.Scale(dist1))

				lV := light.Subtract(intersect).Normalize()
				n := s.Pos.Subtract(intersect).Normalize()
				if n.Z > 0 {
					//fmt.Printf("Negative n.Z\n")
					fmt.Printf("Positive n.Z\n")
				}

				diffLightValue := 0.1 + (diffuseCoefficient * math.Max(0, n.Dot(lV)) * 0.9)
				//fmt.Printf("%s . %s = %.2f\n", lV, n, diffLightValue)
				c := color.RGBA{
					uint8(255 * diffLightValue),
					uint8(255 * diffLightValue),
					uint8(255 * diffLightValue),
					255}

				img.Set(x+255, y+255, c)
			} else {
				img.Set(x+255, y+255, black)
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
