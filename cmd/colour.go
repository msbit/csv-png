package cmd

import (
	"image/color"
	"math"
)

type HSL struct {
	H uint16
	S float64
	L float64
}

func HSLToColour(hsl HSL) color.Color {
	c := (1 - math.Abs((2*hsl.L)-1)) * hsl.S
	x := c * (1 - math.Abs(math.Mod(float64(hsl.H)/60, 2)-1))
	m := hsl.L - c/2

	var r float64
	var g float64
	var b float64
	switch {
	case 0 <= hsl.H && hsl.H < 60:
		r, g, b = c, x, 0.0
	case 60 <= hsl.H && hsl.H < 120:
		r, g, b = x, c, 0.0
	case 120 <= hsl.H && hsl.H < 180:
		r, g, b = 0.0, c, x
	case 180 <= hsl.H && hsl.H < 240:
		r, g, b = 0.0, x, c
	case 240 <= hsl.H && hsl.H < 300:
		r, g, b = x, 0.0, c
	case 300 <= hsl.H && hsl.H < 360:
		r, g, b = c, 0.0, x
	}

	return color.RGBA{uint8((r + m) * 255), uint8((g + m) * 255), uint8((b + m) * 255), 255}
}
