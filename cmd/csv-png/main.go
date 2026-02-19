package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/msbit/csv-png/lib"
)

var (
	input  = flag.String("input", "", "Input CSV file")
	output = flag.String("output", "", "Output PNG file")
	width  = flag.Float64("width", 1920, "Output PNG width")
	height = flag.Float64("height", 1080, "Output PNG height")
)

func main() {
	flag.Parse()

	if *input == "" || *output == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	_, data, err := lib.ReadInput(*input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img := lib.NewImage(*width, *height)
	img.Fill(color.Gray16{0x4000})

	img.DrawAxes()
	img.DrawData(data)

	output, err := os.Create(*output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer output.Close()

	if err := png.Encode(output, img); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
