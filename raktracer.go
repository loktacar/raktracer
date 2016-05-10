package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"
	"os"

	. "github.com/loktacar/raktracer/lib"
)

var shininess = 550.00

type SceneType struct {
	Spheres []SphereType
	Planes  []PlaneType
	Lights  []LightType
	Camera  CameraType
}

type SphereType struct {
	Pos Vector
	R   float64
}

type PlaneType struct {
	Pos  Vector
	Norm Vector
}

type LightType struct {
	Pos       Vector
	Color     ColorType
	Intensity float64
}

type CameraType struct {
	Pos       Vector
	Fov       float64
	ImgWidth  int
	ImgHeight int
}

type ColorType struct {
	R uint8
	G uint8
	B uint8
}

func main() {
	sceneData, err := os.Open("scene.json")
	if err != nil {
		fmt.Printf("fuck")
		return
	}

	sceneParser := json.NewDecoder(sceneData)
	var scene SceneType
	sceneParser.Decode(&scene)

	var spheres []Sphere
	for _, s := range scene.Spheres {
		spheres = append(spheres, NewSphere(s.Pos, s.R))
	}
	var planes []Plane
	for _, p := range scene.Planes {
		planes = append(planes, NewPlane(p.Pos, p.Norm))
	}
	var lights []Light
	for _, l := range scene.Lights {
		lights = append(lights, NewLight(
			l.Pos,
			l.Color.R,
			l.Color.G,
			l.Color.B,
			l.Intensity))
	}

	cam := NewCamera(scene.Camera.Pos, scene.Camera.Fov, scene.Camera.ImgWidth, scene.Camera.ImgHeight)

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
		intersect := r.Pos.Add(r.Dir.Scale(hitDist))
		n := hitSphere.NormalVector(intersect)

		planeCloser := false
		var hitPlane Plane
		for _, p := range planes {
			i, dist := p.Intersects(r)
			if i && (hitDist == -1.00 || dist < hitDist) {
				hitDist = dist
				hitPlane = p
				planeCloser = true
			}
		}
		if planeCloser {
			intersect = r.Pos.Add(r.Dir.Scale(hitDist))
			n = hitPlane.NormalVector(intersect)
		}

		if hitDist == -1.00 {
			continue
		}

		rIntensity := 0.0
		gIntensity := 0.0
		bIntensity := 0.0

		for _, l := range lights {
			lightVector := l.Pos.Subtract(intersect).Normalize()

			lightIntersection := false
			lightRay := Ray{intersect.Add(n.Scale(0.0001)), lightVector}
			for _, s2 := range spheres {
				iL, _ := s2.Intersects(lightRay)
				if iL {
					lightIntersection = true
					break
				}
			}
			if !lightIntersection {
				for _, p2 := range planes {
					iL, _ := p2.Intersects(lightRay)
					if iL {
						lightIntersection = true
						break
					}
				}
			}

			reflectiveVector := n.Scale(2 * lightVector.Scale(-1).Dot(n)).Subtract(lightVector.Scale(-1)).Normalize()

			specularValue := math.Pow(math.Max(0, reflectiveVector.Dot(r.Dir)), shininess)

			diffuseValue := math.Max(0, lightVector.Dot(n))

			// Ambient light 10%
			shadingValue := 0.10

			// Diffuse light 60%
			shadingValue += 0.60 * diffuseValue

			// Specular light 20%
			shadingValue += 0.20 * specularValue

			if lightIntersection {
				shadingValue *= 0.05
			}

			rIntensity += float64(l.C.R) * shadingValue * l.Intensity
			gIntensity += float64(l.C.G) * shadingValue * l.Intensity
			bIntensity += float64(l.C.B) * shadingValue * l.Intensity
		}

		c := color.RGBA{
			uint8(math.Min(255.0, rIntensity)),
			uint8(math.Min(255.0, gIntensity)),
			uint8(math.Min(255.0, bIntensity)),
			255}

		cam.SetPoint(point.ImageX, point.ImageY, c)
	}

	cam.Save("image.png")

	fmt.Println("Done!")
}
