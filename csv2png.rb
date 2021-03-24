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
    ys = row.map(&:to_f)
    data[x] = ys
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
  ymin = 0.0
  ymax = 0.0
  series_count = 0

  data.each do |x, ys|
    xmin = x < xmin ? x : xmin
    xmax = x > xmax ? x : xmax

    series_count = ys.count > series_count ? ys.count : series_count

    ys.each do |y|
      ymin = y < ymin ? y : ymin
      ymax = y > ymax ? y : ymax
    end
  end

  colours = (0...series_count).map { |i| ChunkyPNG::Color.from_hsv((i * 360) / series_count, 1, 1) }

  [colours, scaler(xmin, xmax, MARGIN, WIDTH - MARGIN), scaler(ymin, ymax, HEIGHT - MARGIN, MARGIN)]
end

def draw_data(output, data)
  colours, x_scaler, y_scaler = calculate_attributes(data)
  data.each_cons(2) do |(x0, ys0), (x1, ys1)|
    (0...ys0.count).each do |i|
      output.line_xiaolin_wu(x_scaler.call(x0).to_i, y_scaler.call(ys0[i]).to_i,
                             x_scaler.call(x1).to_i, y_scaler.call(ys1[i]).to_i,
                             colours[i])
    end
  end
end

labels, data = read_input(ARGV[0])

png = ChunkyPNG::Image.new(WIDTH, HEIGHT, ChunkyPNG::Color::WHITE)

draw_axes(png, labels, data)
draw_data(png, data)

png.save(ARGV[1], interlace: true)
