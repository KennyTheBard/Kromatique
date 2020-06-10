package strategy

import (
	"image/color"

	"../utils"
)

// ColorShift is a function that shifts the value of a color by
// a certain degree, be it negative or positive
type ColorShift func(color.Color, int) color.Color

// UniformColorShift modifies all 3 color channels with the entire value
func UniformColorShift(c color.Color, val int) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA64{
		R: uint16(utils.ClampUint8(float64(int(r) + val))),
		G: uint16(utils.ClampUint8(float64(int(g) + val))),
		B: uint16(utils.ClampUint8(float64(int(b) + val))),
		A: uint16(a),
	}
}

// DividedColorShift modifies all 3 color channels by equal part
// of shifting value, clamping the result
func DividedColorShift(c color.Color, val int) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA64{
		R: uint16(utils.ClampUint8(float64(int(r) + val/3))),
		G: uint16(utils.ClampUint8(float64(int(g) + val/3))),
		B: uint16(utils.ClampUint8(float64(int(b) + val/3))),
		A: uint16(a),
	}
}
