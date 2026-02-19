package lib

import (
	"image"
	"image/color"
	"maps"
	"math"
	"slices"
)

type Image struct {
	*image.RGBA
	width  float64
	height float64
	margin float64
}

type point struct {
	x float64
	y float64
}

func NewImage(width, height float64) Image {
	return Image{
		image.NewRGBA(
			image.Rectangle{
				image.Point{0, 0},
				image.Point{int(width), int(height)},
			},
		),
		width,
		height,
		margin(width, height),
	}
}

func margin(width, height float64) float64 {
	return min(width, height) / 20.0
}

func (img *Image) DrawLine(from point, to point, hsl hsl) {
	steep := math.Abs(to.y-from.y) > math.Abs(to.x-from.x)
	if steep {
		from.x, from.y = from.y, from.x
		to.x, to.y = to.y, to.x
	}

	if from.x > to.x {
		from.x, to.x = to.x, from.x
		from.y, to.y = to.y, from.y
	}

	dx := to.x - from.x
	dy := to.y - from.y

	gradient := dy / dx
	if dx == 0.0 {
		gradient = 1.0
	}

	xend := math.Round(from.x)
	yend := math.Round(from.y + gradient*(xend-from.x))
	xgap := rfpart(from.x + 0.5)
	xpxl1 := xend
	ypxl1 := math.Floor(yend)
	if steep {
		img.plot(ypxl1, xpxl1, rfpart(yend)*xgap, hsl)
		img.plot(ypxl1+1, xpxl1, fpart(yend)*xgap, hsl)
	} else {
		img.plot(xpxl1, ypxl1, rfpart(yend)*xgap, hsl)
		img.plot(xpxl1, ypxl1+1, fpart(yend)*xgap, hsl)
	}
	intery := yend + gradient

	xend = math.Round(to.x)
	yend = to.y + gradient*(xend-to.x)
	xgap = fpart(to.x + 0.5)
	xpxl2 := xend
	ypxl2 := math.Floor(yend)
	if steep {
		img.plot(ypxl2, xpxl2, rfpart(yend)*xgap, hsl)
		img.plot(ypxl2+1, xpxl2, fpart(yend)*xgap, hsl)
	} else {
		img.plot(xpxl2, ypxl2, rfpart(yend)*xgap, hsl)
		img.plot(xpxl2, ypxl2+1, fpart(yend)*xgap, hsl)
	}

	if steep {
		for x := xpxl1 + 1; x < xpxl2; x++ {
			img.plot(math.Floor(intery), x, rfpart(intery), hsl)
			img.plot(math.Floor(intery)+1, x, fpart(intery), hsl)
			intery = intery + gradient
		}
	} else {
		for x := xpxl1 + 1; x < xpxl2; x++ {
			img.plot(x, math.Floor(intery), rfpart(intery), hsl)
			img.plot(x, math.Floor(intery)+1, fpart(intery), hsl)
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

func (img *Image) plot(x, y, brightness float64, hsl hsl) {
	hsl.l = 1.0 - (brightness * 0.5)
	img.Set(int(x), int(y), hSLToColour(hsl))
}

func (i *Image) DrawAxes() {
	i.DrawLine(
		point{i.margin, i.margin},
		point{i.margin, i.height - i.margin},
		hsl{0, 0.0, 0.0},
	)
	i.DrawLine(
		point{i.margin, i.height - i.margin},
		point{i.width - i.margin, i.height - i.margin},
		hsl{0, 0.0, 0.0},
	)
}

func (img *Image) DrawData(data map[float64][]float64) {
	colours, hScaler, vScaler := calculateAttributes(img, data)
	keys := slices.Sorted(maps.Keys(data))
	for i := 1; i < len(keys); i++ {
		x0 := keys[i-1]
		x1 := keys[i]
		series0 := data[x0]
		series1 := data[x1]
		for j := 0; j < len(series0); j++ {
			img.DrawLine(
				point{hScaler(x0), vScaler(series0[j])},
				point{hScaler(x1), vScaler(series1[j])},
				colours[j],
			)
		}
	}
}

func calculateAttributes(
	img *Image,
	data map[float64][]float64,
) ([]hsl, scaler, scaler) {
	valueMin := math.Inf(1)
	valueMax := math.Inf(-1)
	seriesCount := 0

	keys := slices.Collect(maps.Keys(data))
	values := slices.Collect(maps.Values(data))

	xmin := slices.Min(keys)
	xmax := slices.Max(keys)

	for _, series := range values {
		seriesCount = max(seriesCount, len(series))

		valueMin = min(valueMin, slices.Min(series))
		valueMax = max(valueMax, slices.Max(series))
	}

	colours := make([]hsl, seriesCount)
	for i := 0; i < seriesCount; i++ {
		colours[i] = hsl{uint16((360 * i) / seriesCount), 1.0, 0.5}
	}

	return colours,
		Scaler(xmin, xmax, img.margin, img.width-img.margin),
		Scaler(valueMin, valueMax, img.height-img.margin, img.margin)
}
