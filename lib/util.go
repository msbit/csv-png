package lib

type scaler = func(float64) float64

func Scaler(inputMin, inputMax, outputMin, outputMax float64) scaler {
	scale := (outputMax - outputMin) / (inputMax - inputMin)
	return func(input float64) float64 {
		return ((input - inputMin) * scale) + outputMin
	}
}
