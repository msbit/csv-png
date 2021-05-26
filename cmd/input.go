package cmd

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func ReadInput(input string) ([]string, map[float64][]float64, error) {
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
