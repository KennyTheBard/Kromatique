package histogram

import (
	"../utils"
	"image"
	"image/color"
	"math"
)

// ColorShift is a function that shifts the value of a color by
// a certain degree, be it negative or positive
type ColorShift func(color.Color, int) color.Color

// UniformColorShift modifies all 3 color channels with the entire value
func UniformColorShift(c color.Color, val int) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(utils.ClampUint8(float64(int(r>>8) + val))),
		G: uint8(utils.ClampUint8(float64(int(g>>8) + val))),
		B: uint8(utils.ClampUint8(float64(int(b>>8) + val))),
		A: uint8(a >> 8),
	}
}

// DividedColorShift modifies all 3 color channels by equal part
// of shifting value, clamping the result
func DividedColorShift(c color.Color, val int) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(utils.ClampUint8(float64(int(r>>8) + val/3))),
		G: uint8(utils.ClampUint8(float64(int(g>>8) + val/3))),
		B: uint8(utils.ClampUint8(float64(int(b>>8) + val/3))),
		A: uint8(a >> 8),
	}
}

func Equalize(img image.Image, histogram *Histogram, shift ColorShift) image.Image {
	values := histogram.Values()
	cumulative := histogram.Cumulative()
	idealCumulative := make([]uint, len(values))

	// calculate the length of the value space
	space := uint(len(values))

	// calculate slope of ideal cumulative
	slope := float64(cumulative[len(cumulative)-1]) / float64(space)

	// create a mapper for each interval
	for idx, _ := range values {
		idealCumulative[idx] = uint(math.Round(slope * float64(idx)))
	}

	return Matching(img, histogram, idealCumulative, shift)
}
