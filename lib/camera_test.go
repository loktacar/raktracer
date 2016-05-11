package raktracer

import (
	"image"
	"testing"
)

func TestCameraConstructor(t *testing.T) {
	cases := []struct {
		o   Vector
		fov float64
		iW  int
		iH  int
	}{
		{Vector{0, 0, 0}, 90, 10, 5},
	}
	for _, c := range cases {
		got := NewCamera(c.o, c.fov, c.iW, c.iH)
		if got.origin != c.o {
			t.Errorf("camera.origin = %s, want %s", got.origin, c.o)
		}
		if got.fov != c.fov {
			t.Errorf("camera.fov = %.0f, want %.0f", got.fov, c.fov)
		}
		if got.imgWidth != c.iW {
			t.Errorf("camera.imgWidth = %d, want %d", got.imgWidth, c.iW)
		}
		if got.imgHeight != c.iH {
			t.Errorf("camera.imgHeight = %d, want %d", got.imgHeight, c.iH)
		}
		if got.img.Rect.Min.X != 0 || got.img.Rect.Min.Y != 0 ||
			got.img.Rect.Max.X != c.iW || got.img.Rect.Max.Y != c.iH {
			t.Errorf("camera.img.Rect = %s, want %s", got.img.Rect, image.Rect(0, 0, c.iW, c.iH))
		}
	}
}

func TestCameraImagePoints(t *testing.T) {
	width, height := 2, 2
	c := NewCamera(Vector{0, 0, 0}, 90, width, height)

	want := []struct {
		x int
		y int
		r Ray
	}{
		{0, 0, Ray{Vector{0, 0, 0}, Vector{-0.408248, 0.408248, -0.816497}}},
		{0, 1, Ray{Vector{0, 0, 0}, Vector{-0.408248, -0.408248, -0.816497}}},
		{1, 0, Ray{Vector{0, 0, 0}, Vector{0.408248, 0.408248, -0.816497}}},
		{1, 1, Ray{Vector{0, 0, 0}, Vector{0.408248, -0.408248, -0.816497}}},
	}

	i := 0
	for ip := range c.ImagePoints() {
		if ip.ImageX != want[i].x || ip.ImageY != want[i].y ||
			ip.SceneRay != want[i].r {
			t.Errorf("ip[%d] = %d, %d, %s; want %d, %d, %s", i, ip.ImageX, ip.ImageY, ip.SceneRay, want[i].x, want[i].y, want[i].r)
			t.Errorf("%f %f %f", ip.SceneRay.Dir.X, ip.SceneRay.Dir.Y, ip.SceneRay.Dir.Z)
		}

		i += 1
	}
}
