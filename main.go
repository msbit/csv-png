package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"encoding/csv"
)

type options_t struct {
	input  string
	output string
	width  uint
	height uint
	margin uint
}

func main() {
	options := options_t{margin: 54}

	flag.StringVar(&options.input, "input", "", "Input CSV file")
	flag.StringVar(&options.output, "output", "", "Output PNG file")

	flag.UintVar(&options.width, "width", 1920, "Output PNG width")
	flag.UintVar(&options.height, "height", 1080, "Output PNG height")

	flag.Parse()

	if options.input == "" || options.output == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	labels, data, err := read_input(options.input)
	if err != nil {
		os.Exit(1)
	}

	/*
	   TODO:
	     * calculate attributes
	     * draw the data
	*/

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
