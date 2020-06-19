package strategy

import (
	"../utils"
	"image/color"
)

// ColorCorrection is a function that shifts the value of a color by
// a certain degree, be it negative or positive
type ColorCorrection func(color.Color, int16) color.Color

// UniformColorCorrection modifies all 3 color channels with the entire value
func UniformColorCorrection(c color.Color, val int16) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(utils.ClampUint8(float64(r>>8) + float64(val))),
		G: uint8(utils.ClampUint8(float64(g>>8) + float64(val))),
		B: uint8(utils.ClampUint8(float64(b>>8) + float64(val))),
		A: uint8(a >> 8),
	}
}

// DividedColorCorrection modifies all 3 color channels by equal part
// of shifting value, clamping the result
func DividedColorCorrection(c color.Color, val int16) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(utils.ClampUint8(float64(r>>8) + float64(val)/3)),
		G: uint8(utils.ClampUint8(float64(g>>8) + float64(val)/3)),
		B: uint8(utils.ClampUint8(float64(b>>8) + float64(val)/3)),
		A: uint8(a >> 8),
	}
}

func LightnessCorrection(c color.Color, val int16) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(utils.ClampUint8(float64(r>>8) + float64(val)*0.2126)),
		G: uint8(utils.ClampUint8(float64(g>>8) + float64(val)*0.7152)),
		B: uint8(utils.ClampUint8(float64(b>>8) + float64(val)*0.0722)),
		A: uint8(a >> 8),
	}
}
