package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strconv"
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

	labels, data, err := read_input(options.input)
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

func read_input(input string) ([]string, map[float64][]float64, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, nil, err
	}

	defer f.Close()

	reader := csv.NewReader(f)

	labels, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return nil, nil, nil
		}

		return nil, nil, err
	}

	data := map[float64][]float64{}

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				return labels, data, nil
			}

			return nil, nil, err
		}

		x, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			return nil, nil, err
		}

		series := []float64{}
		for _, value := range row[1:] {
			y, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, nil, err
			}
			series = append(series, y)
		}
		data[x] = series
	}
}
