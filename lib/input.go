package lib

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

	var labels []string

	data := map[float64][]float64{}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			return labels, data, nil
		}

		if err != nil {
			return nil, nil, err
		}

		if labels == nil {
			labels = row
			continue
		}

		x, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			return nil, nil, err
		}

		ys := []float64{}
		for _, value := range row[1:] {
			y, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, nil, err
			}
			ys = append(ys, y)
		}
		data[x] = ys
	}
}
