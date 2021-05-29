package cmd

import (
	"image"
	"image/color"
	"math"
)

type Image struct {
	*image.RGBA
}

func (img *Image) DrawLine(x0 float64, y0 float64, x1 float64, y1 float64, c color.Color) {
	steep := math.Abs(y1-y0) > math.Abs(x1-x0)
	if steep {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}

	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	dx := x1 - x0
	dy := y1 - y0

	gradient := dy / dx
	if dx == 0.0 {
		gradient = 1.0
	}

	xend := math.Round(x0)
	yend := math.Round(y0 + gradient*(xend-x0))
	xgap := rfpart(x0 + 0.5)
	xpxl1 := xend
	ypxl1 := math.Floor(yend)
	if steep {
		img.plot(ypxl1, xpxl1, rfpart(yend)*xgap, c)
		img.plot(ypxl1+1, xpxl1, fpart(yend)*xgap, c)
	} else {
		img.plot(xpxl1, ypxl1, rfpart(yend)*xgap, c)
		img.plot(xpxl1, ypxl1+1, fpart(yend)*xgap, c)
	}
	intery := yend + gradient

	xend = math.Round(x1)
	yend = y1 + gradient*(xend-x1)
	xgap = fpart(x1 + 0.5)
	xpxl2 := xend
	ypxl2 := math.Floor(yend)
	if steep {
		img.plot(ypxl2, xpxl2, rfpart(yend)*xgap, c)
		img.plot(ypxl2+1, xpxl2, fpart(yend)*xgap, c)
	} else {
		img.plot(xpxl2, ypxl2, rfpart(yend)*xgap, c)
		img.plot(xpxl2, ypxl2+1, fpart(yend)*xgap, c)
	}

	if steep {
		for x := xpxl1 + 1; x < xpxl2; x++ {
			img.plot(math.Floor(intery), x, rfpart(intery), c)
			img.plot(math.Floor(intery)+1, x, fpart(intery), c)
			intery = intery + gradient
		}
	} else {
		for x := xpxl1 + 1; x < xpxl2; x++ {
			img.plot(x, math.Floor(intery), rfpart(intery), c)
			img.plot(x, math.Floor(intery)+1, fpart(intery), c)
			intery = intery + gradient
		}
	}
}

func (img *Image) Fill(c color.Color) {
	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			img.Set(x, y, c)
		}
	}
}

func rfpart(x float64) float64 {
	return 1 - fpart(x)
}

func fpart(x float64) float64 {
	return x - math.Floor(x)
}

func (img *Image) plot(x float64, y float64, brightness float64, full color.Color) {
	r, g, b, _ := full.RGBA()
	c := color.RGBA{
		uint8(float64(r) * brightness / 256.0),
		uint8(float64(g) * brightness / 256.0),
		uint8(float64(b) * brightness / 256.0),
		uint8(brightness * 255.0)}
	img.Set(int(x), int(y), c)
}
