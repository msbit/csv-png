package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"sort"

	_color "github.com/gerow/go-color"

	"github.com/msbit/csv-png/cmd"
)

type scaler = func(float64) float64

type options_t struct {
	input  string
	output string
	width  int
	height int
	margin int
}

func main() {
	options := options_t{margin: 54}

	flag.StringVar(&options.input, "input", "", "Input CSV file")
	flag.StringVar(&options.output, "output", "", "Output PNG file")

	flag.IntVar(&options.width, "width", 1920, "Output PNG width")
	flag.IntVar(&options.height, "height", 1080, "Output PNG height")

	flag.Parse()

	if options.input == "" || options.output == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	_, data, err := cmd.ReadInput(options.input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{options.width, options.height}})
	cmd.Fill(img, color.White)

	draw_axes(img, options)
	draw_data(img, data, options)

	output, err := os.Create(options.output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer output.Close()

	err = png.Encode(output, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func draw_axes(img *image.RGBA, options options_t) {
	margin := float64(options.margin)
	width := float64(options.width)
	height := float64(options.height)

	cmd.DrawLine(img, margin, margin, margin, height-margin, color.Black)
	cmd.DrawLine(img, margin, height-margin, width-margin, height-margin, color.Black)
}

func draw_data(img *image.RGBA, data map[float64][]float64, options options_t) {
	colours, h_scaler, v_scaler := calculate_attributes(data, options)
	keys := make([]float64, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	for i := 1; i < len(keys); i++ {
		x0 := keys[i-1]
		x1 := keys[i]
		series0 := data[x0]
		series1 := data[x1]
		for j := 0; j < len(series0); j++ {
			cmd.DrawLine(img, h_scaler(x0), v_scaler(series0[j]), h_scaler(x1), v_scaler(series1[j]), colours[j])
		}
	}
}

func calculate_attributes(data map[float64][]float64, options options_t) ([]color.Color, scaler, scaler) {
	xmin := math.Inf(1)
	xmax := math.Inf(-1)
	value_min := math.Inf(1)
	value_max := math.Inf(-1)
	series_count := 0

	for x, series := range data {
		xmin = math.Min(x, xmin)
		xmax = math.Max(x, xmax)
		series_count = cmd.Max(series_count, len(series))

		for _, value := range series {
			value_min = math.Min(value_min, value)
			value_max = math.Max(value_max, value)
		}
	}

	colours := []color.Color{}
	for i := 0; i < series_count; i++ {
		hue := float64(i) / float64(series_count)
		rgb := _color.HSL{hue, 0.5, 0.5}.ToRGB()
		colours = append(colours, color.RGBA{uint8(rgb.R * 0xff), uint8(rgb.G * 0xff), uint8(rgb.B * 0xff), 0xff})
	}

	margin := options.margin
	width := options.width
	height := options.height

	return colours,
		cmd.Scaler(xmin, xmax, float64(margin), float64(width-margin)),
		cmd.Scaler(value_min, value_max, float64(height-margin), float64(margin))
}
