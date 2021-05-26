package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/msbit/csv-png/cmd"
)

type options_t struct {
	input  string
	output string
	width  int
	height int
	margin uint
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

	labels, data, err := cmd.ReadInput(options.input)
	if err != nil {
		os.Exit(1)
	}

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{options.width, options.height}})
	for x := 0; x < options.width; x++ {
		for y := 0; y < options.height; y++ {
			img.Set(x, y, color.White)
		}
	}

	/*
	   TODO:
	     * calculate attributes
	     * draw the data
	*/

	output, err := os.Create(options.output)
	if err != nil {
		os.Exit(1)
	}

	defer output.Close()

	png.Encode(output, img)

	fmt.Println(labels, data)
}
