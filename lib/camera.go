package raktracer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Camera struct {
	//TopLeft           Vector
	//BottomRight       Vector
	Origin            Vector
	OutputImageWidth  int
	OutputImageHeight int
	img               *image.RGBA
}
type CameraPoint struct {
	ImageX   int
	ImageY   int
	ScenePos Vector
}

func NewCamera(origin Vector, width int, height int) Camera {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}

	return Camera{
		origin,
		width,
		height,
		img,
	}
}

func (c Camera) ImagePoints() chan CameraPoint {
	ch := make(chan CameraPoint)

	go func() {
		for x := -c.OutputImageWidth / 2; x < c.OutputImageWidth/2; x++ {
			for y := -c.OutputImageHeight / 2; y < c.OutputImageHeight/2; y++ {
				ch <- CameraPoint{x, y, Vector{float64(x), float64(y), 0}}
			}
		}

		close(ch)
	}()

	return ch
}

func (c Camera) CameraRay(pos Vector) Ray {
	return Ray{c.Origin, pos.Subtract(c.Origin).Normalize()}
}

func (c Camera) SetPoint(x int, y int, color color.RGBA) {
	fmt.Printf("x: %d, y: %d, color: %s\n", x, y, color)
	c.img.Set(x+c.OutputImageWidth/2, y+c.OutputImageHeight/2, color)
}

func (c Camera) Save(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, c.img)
}
