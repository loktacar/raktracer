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

var ambientCoefficient = 0.01
var shadingCoefficient = 0.05

type SceneType struct {
	Spheres []SphereType
	Planes  []PlaneType
	Lights  []LightType
	Camera  CameraType
}

type SphereType struct {
	Pos                 Vector
	R                   float64
	DiffuseCoefficient  float64
	SpecularCoefficient float64
	SpecularN           float64
}

type PlaneType struct {
	Pos                 Vector
	Norm                Vector
	DiffuseCoefficient  float64
	SpecularCoefficient float64
	SpecularN           float64
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

	var objects []Object
	for _, s := range scene.Spheres {
		objects = append(objects, NewSphere(s.Pos, s.R, s.DiffuseCoefficient, s.SpecularCoefficient, s.SpecularN, 0.0))
	}
	for _, p := range scene.Planes {
		objects = append(objects, NewPlane(p.Pos, p.Norm, p.DiffuseCoefficient, p.SpecularCoefficient, p.SpecularN, 0.0))
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
		//var hitObject Object
		var intersect Vector
		var n Vector
		//var refDir Vector
		var dC float64
		var sC float64
		var sN float64
		//var rC float64
		for _, o := range objects {
			i, dist := o.Intersects(r)
			if i && (hitDist == -1 || dist < hitDist) {
				//hitObject = o
				hitDist = dist
				intersect = r.Pos.Add(r.Dir.Scale(hitDist))
				n, _, dC, sC, sN, _ = o.SurfaceProperties(intersect, r.Dir)
			}
		}

		if hitDist == -1.00 {
			continue
		}

		rIntensity := ambientCoefficient * 255
		gIntensity := ambientCoefficient * 255
		bIntensity := ambientCoefficient * 255

		for _, l := range lights {
			lightVector := l.Pos.Subtract(intersect)
			lightDistance := lightVector.Length()
			lightVector = lightVector.Normalize()

			lightIntersection := false
			lightRay := Ray{intersect.Add(n.Scale(0.0001)), lightVector}
			for _, o2 := range objects {
				iL, t := o2.Intersects(lightRay)
				if iL && t <= lightDistance {
					lightIntersection = true
					break
				}
			}

			reflectiveVector := n.Scale(2 * lightVector.Scale(-1).Dot(n)).Subtract(lightVector.Scale(-1)).Normalize()

			specularValue := math.Pow(math.Max(0, reflectiveVector.Dot(r.Dir)), sN)
			diffuseValue := math.Max(0, lightVector.Dot(n))

			shadingValue := dC*diffuseValue + sC*specularValue

			if lightIntersection {
				shadingValue *= shadingCoefficient
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
