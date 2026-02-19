package lib

type scalerFunc = func(float64) float64

func scaler(inputMin, inputMax, outputMin, outputMax float64) scalerFunc {
	scale := (outputMax - outputMin) / (inputMax - inputMin)
	return func(input float64) float64 {
		return ((input - inputMin) * scale) + outputMin
	}
}
