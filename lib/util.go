package lib

type scaler = func(float64) float64

type Options struct {
	Input  string
	Output string
	Width  int
	Height int
	Margin int
}

func Scaler(inputMin float64, inputMax float64, outputMin float64, outputMax float64) scaler {
	scale := (outputMax - outputMin) / (inputMax - inputMin)
	return func(input float64) float64 {
		return ((input - inputMin) * scale) + outputMin
	}
}
