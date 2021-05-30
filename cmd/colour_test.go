package cmd

import (
	"image/color"
	"testing"
)

var testCases = map[HSL]color.RGBA{
	/* Black */ HSL{0, 0.0, 0.0}: color.RGBA{0, 0, 0, 255},
	/* White */ HSL{0, 0.0, 1.0}: color.RGBA{255, 255, 255, 255},
	/* Red */ HSL{0, 1.0, 0.5}: color.RGBA{255, 0, 0, 255},
	/* Lime */ HSL{120, 1.0, 0.5}: color.RGBA{0, 255, 0, 255},
	/* Blue */ HSL{240, 1.0, 0.5}: color.RGBA{0, 0, 255, 255},
	/* Yellow */ HSL{60, 1.0, 0.5}: color.RGBA{255, 255, 0, 255},
	/* Cyan */ HSL{180, 1.0, 0.5}: color.RGBA{0, 255, 255, 255},
	/* Magenta */ HSL{300, 1.0, 0.5}: color.RGBA{255, 0, 255, 255},
	/* Silver */ HSL{0, 0.0, 0.75}: color.RGBA{191, 191, 191, 255},
	/* Gray */ HSL{0, 0.0, 0.5}: color.RGBA{127, 127, 127, 255},
	/* Maroon */ HSL{0, 1.0, 0.25}: color.RGBA{127, 0, 0, 255},
	/* Olive */ HSL{60, 1.0, 0.25}: color.RGBA{127, 127, 0, 255},
	/* Green */ HSL{120, 1.0, 0.25}: color.RGBA{0, 127, 0, 255},
	/* Purple */ HSL{300, 1.0, 0.25}: color.RGBA{127, 0, 127, 255},
	/* Teal */ HSL{180, 1.0, 0.25}: color.RGBA{0, 127, 127, 255},
	/* Navy */ HSL{240, 1.0, 0.25}: color.RGBA{0, 0, 127, 255},
}

func TestHSLToColour(t *testing.T) {
	for input, want := range testCases {
		got := HSLToColour(input)
		if got != want {
			r, g, b, a := got.RGBA()
			t.Fatalf("input: %v, got {%d %d %d %d}, wanted {%d %d %d %d}", input, r, g, b, a, want.R, want.G, want.B, want.A)
		}
	}
}
