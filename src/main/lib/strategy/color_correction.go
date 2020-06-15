package strategy

import (
	"image/color"
	"math"

	"../utils"
)

// ColorCorrection is a function that shifts the value of a color by
// a certain degree, be it negative or positive
type ColorCorrection func(color.Color, float64) color.Color

// UniformColorCorrection modifies all 3 color channels with the entire value
func UniformColorCorrection(c color.Color, val float64) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA64{
		R: uint16(utils.ClampUint16(float64(r) + val*math.MaxUint16)),
		G: uint16(utils.ClampUint16(float64(g) + val*math.MaxUint16)),
		B: uint16(utils.ClampUint16(float64(b) + val*math.MaxUint16)),
		A: uint16(a),
	}
}

// DividedColorCorrection modifies all 3 color channels by equal part
// of shifting value, clamping the result
func DividedColorCorrection(c color.Color, val float64) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA64{
		R: uint16(utils.ClampUint16(float64(r) + val*math.MaxUint16/3)),
		G: uint16(utils.ClampUint16(float64(g) + val*math.MaxUint16/3)),
		B: uint16(utils.ClampUint16(float64(b) + val*math.MaxUint16/3)),
		A: uint16(a),
	}
}

func LightnessCorrection(c color.Color, val float64) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA64{
		R: uint16(utils.ClampUint16(float64(r) + val*math.MaxUint16*0.2126)),
		G: uint16(utils.ClampUint16(float64(g) + val*math.MaxUint16*0.7152)),
		B: uint16(utils.ClampUint16(float64(b) + val*math.MaxUint16*0.0722)),
		A: uint16(a),
	}
}
