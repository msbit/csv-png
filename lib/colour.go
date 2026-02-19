package lib

import (
	"image/color"
	"math"
)

type hsl struct {
	h uint16
	s float64
	l float64
}

func hSLToColour(hsl hsl) color.Color {
	c := (1 - math.Abs((2*hsl.l)-1)) * hsl.s
	x := c * (1 - math.Abs(math.Mod(float64(hsl.h)/60, 2)-1))
	m := hsl.l - c/2

	var r float64
	var g float64
	var b float64
	switch {
	case 0 <= hsl.h && hsl.h < 60:
		r, g, b = c, x, 0.0
	case 60 <= hsl.h && hsl.h < 120:
		r, g, b = x, c, 0.0
	case 120 <= hsl.h && hsl.h < 180:
		r, g, b = 0.0, c, x
	case 180 <= hsl.h && hsl.h < 240:
		r, g, b = 0.0, x, c
	case 240 <= hsl.h && hsl.h < 300:
		r, g, b = x, 0.0, c
	case 300 <= hsl.h && hsl.h < 360:
		r, g, b = c, 0.0, x
	}

	return color.RGBA{
		uint8((r + m) * 255),
		uint8((g + m) * 255),
		uint8((b + m) * 255),
		255,
	}
}
