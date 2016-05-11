package raktracer

import (
	"fmt"
	"image/color"
)

type Light struct {
	Pos       Vector
	C         color.RGBA
	Intensity float64
}

func NewLight(pos Vector, r uint8, g uint8, b uint8, intensity float64) Light {
	return Light{
		pos,
		color.RGBA{
			r,
			g,
			b,
			255,
		},
		intensity,
	}
}

func (l Light) String() string {
	return fmt.Sprintf("Light{Pos: %s C: %s Intensity: %.0f}", l.Pos, l.C, l.Intensity)
}
