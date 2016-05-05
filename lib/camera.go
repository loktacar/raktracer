package raktracer

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

// A Camera represents a camera in the scene, it calculates the rays which
// will be cast to create the final image, and collectes the colours each ray
// generates into the final image.
type Camera struct {
	origin    Vector
	fov       float64
	imgWidth  int
	imgHeight int
	img       *image.RGBA
}

// An ImagePointRealRay represents a point on the image and it's corresponding camera ray, from the camera origin and through the images x,y pixel in the scene space.
type ImagePointRealRay struct {
	ImageX   int
	ImageY   int
	SceneRay Ray
}

// NewCamera returns a new camera with the given origin point, field-of-view and image height and width.
func NewCamera(origin Vector, fov float64, width int, height int) Camera {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}

	return Camera{
		origin,
		fov,
		width,
		height,
		img,
	}
}

// ImagePoints returns a channel of each x,y pixel on the cameras image and it's corresponding camera ray.
func (c Camera) ImagePoints() chan ImagePointRealRay {
	// This algorithm based (stolen) from here http://www.scratchapixel.com/lessons/3d-basic-rendering/ray-tracing-generating-camera-rays/generating-camera-rays
	ch := make(chan ImagePointRealRay)

	floatWidth := float64(c.imgWidth)
	floatHeight := float64(c.imgHeight)

	widthAspectRatio := float64(1)
	heightAspectRatio := float64(1)
	if c.imgWidth > c.imgHeight {
		widthAspectRatio = floatWidth / floatHeight
	} else {
		heightAspectRatio = floatHeight / floatWidth
	}

	fovRatio := math.Tan(c.fov / 2 * math.Pi / 180)

	go func() {
		for x := float64(0); x < floatWidth; x++ {
			for y := float64(0); y < floatHeight; y++ {
				dir := Vector{
					(2*((x+0.5)/floatWidth) - 1) * fovRatio * widthAspectRatio,
					(1 - 2*(y+0.5)/floatHeight) * fovRatio * heightAspectRatio,
					-1,
				}

				// TODO: matrix multiplication to rotate and translate the camera.

				ch <- ImagePointRealRay{
					int(x),
					int(y),
					Ray{
						c.origin,
						dir.Normalize(),
					},
				}
			}
		}

		close(ch)
	}()

	return ch
}

// Sets a pixel in the cameras image to the given color.
func (c Camera) SetPoint(x int, y int, color color.RGBA) {
	c.img.Set(x, y, color)
}

// Saves the cameras image as a file at the given filename
func (c Camera) Save(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, c.img)
}
