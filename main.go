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

func main() {
	options := cmd.Options{Margin: 54}

	flag.StringVar(&options.Input, "input", "", "Input CSV file")
	flag.StringVar(&options.Output, "output", "", "Output PNG file")

	flag.IntVar(&options.Width, "width", 1920, "Output PNG width")
	flag.IntVar(&options.Height, "height", 1080, "Output PNG height")

	flag.Parse()

	if options.Input == "" || options.Output == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	_, data, err := cmd.ReadInput(options.Input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img := &cmd.Image{
		image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{options.Width, options.Height}}),
	}
	img.Fill(color.White)

	img.DrawAxes(options)
	img.DrawData(data, options)

	output, err := os.Create(options.Output)
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
