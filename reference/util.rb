# frozen_string_literal: true

def scaler(input_min, input_max, output_min, output_max)
  scale = (output_max - output_min).to_f / (input_max - input_min)
  lambda { |input|
    ((input - input_min) * scale) + output_min
  }
end
