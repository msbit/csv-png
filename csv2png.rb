# frozen_string_literal: true

require 'bundler'
require 'csv'

require './util'

Bundler.require(:default)

return if ARGV.count < 2

MARGIN = 50
WIDTH = 1920
HEIGHT = 1080

def read_input(filename)
  input = CSV.foreach(filename)
  data = {}
  labels = []

  input.each.with_index do |row, i|
    if i.zero?
      labels = row
      next
    end

    x = row.shift.to_f
    values = row.map(&:to_f)
    data[x] = values
  end

  [labels, data]
end

def draw_axes(output, _labels, _data)
  output.line_xiaolin_wu(MARGIN, MARGIN, MARGIN, HEIGHT - MARGIN, ChunkyPNG::Color::BLACK)
  output.line_xiaolin_wu(MARGIN, HEIGHT - MARGIN, WIDTH - MARGIN, HEIGHT - MARGIN, ChunkyPNG::Color::BLACK)
end

def calculate_attributes(data)
  xmin = 0.0
  xmax = 0.0
  value_min = 0.0
  value_max = 0.0
  series_count = 0

  data.each do |x, series|
    xmin = x < xmin ? x : xmin
    xmax = x > xmax ? x : xmax

    series_count = series.count > series_count ? series.count : series_count

    series.each do |value|
      value_min = value < value_min ? value : value_min
      value_max = value > value_max ? value : value_max
    end
  end

  colours = (0...series_count).map { |i| ChunkyPNG::Color.from_hsv((i * 360) / series_count, 1, 1) }

  [colours, scaler(xmin, xmax, MARGIN, WIDTH - MARGIN), scaler(value_min, value_max, HEIGHT - MARGIN, MARGIN)]
end

def draw_data(output, data)
  colours, horizontal_scaler, vertical_scaler = calculate_attributes(data)
  data.each_cons(2) do |(x0, series0), (x1, series1)|
    (0...series0.count).each do |i|
      output.line_xiaolin_wu(horizontal_scaler.call(x0).to_i, vertical_scaler.call(series0[i]).to_i,
                             horizontal_scaler.call(x1).to_i, vertical_scaler.call(series1[i]).to_i,
                             colours[i])
    end
  end
end

labels, data = read_input(ARGV[0])

png = ChunkyPNG::Image.new(WIDTH, HEIGHT, ChunkyPNG::Color::WHITE)

draw_axes(png, labels, data)
draw_data(png, data)

png.save(ARGV[1], interlace: true)
