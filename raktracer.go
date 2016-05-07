package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"
	"os"

	. "github.com/loktacar/raktracer/lib"
)

var shininess = 100.00

type SceneType struct {
	Spheres []SphereType
	Planes  []PlaneType
	Light   LightType
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
	Pos Vector
}

type CameraType struct {
	Pos       Vector
	Fov       float64
	ImgWidth  int
	ImgHeight int
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
	light := scene.Light.Pos

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

		lightVector := light.Subtract(intersect).Normalize()

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
