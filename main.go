package main

import (
	"flag"
	"fmt"
	"os"
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

	/*
	   TODO:
	     * read CSV
	     * parse into labels and data
	     * calculate attributes
	     * draw the data
	*/

	fmt.Println(options)
}
