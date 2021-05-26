package cmd

type scaler = func(float64) float64

func Scaler(input_min float64, input_max float64, output_min float64, output_max float64) scaler {
	scale := (output_max - output_min) / (input_max - input_min)
	return func(input float64) float64 {
		return ((input - input_min) * scale) + output_min
	}
}
