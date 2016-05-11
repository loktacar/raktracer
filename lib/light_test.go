package raktracer

import (
	"image/color"
	"testing"
)

func TestLightConstructor(t *testing.T) {
	cases := []struct {
		p    Vector
		r    uint8
		g    uint8
		b    uint8
		i    float64
		want Light
	}{
		{Vector{0, 0, 0}, 0, 0, 0, 0, Light{Vector{0, 0, 0}, color.RGBA{0, 0, 0, 255}, 0}},
	}
	for _, c := range cases {
		got := NewLight(c.p, c.r, c.g, c.b, c.i)
		if got != c.want {
			t.Errorf("LightNew(%s, %s, %s, %s, %s) = %s, want %s", c.p, c.r, c.g, c.b, c.i, got, c.want)
		}
	}
}

func TestLightString(t *testing.T) {
	cases := []struct {
		l    Light
		want string
	}{
		{Light{Vector{0, 0, 0}, color.RGBA{0, 0, 0, 0}, 0}, "Light{Pos:[0.00 0.00 0.00] C:{R:0 G:0 B:0 A:0} Intensity:0}"},
	}
	for _, c := range cases {
		got := c.l.String()
		if got != c.want {
			t.Errorf("%s.String() = %s, want %s", c.l, got, c.want)
		}
	}
}
