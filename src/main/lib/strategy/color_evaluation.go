package strategy

import (
	"image/color"
	"math"

	"../utils"
)

// ColorEvaluation is a function that transform a given color with multiple
// color channels into a single numerical value, easier to index
type ColorEvaluation func(color.Color) uint8

func LightnessEval(c color.Color) uint8 {
	r, g, b, _ := c.RGBA()

	lightness := float64(r)*0.2126 + float64(g)*0.7152 + float64(b)*0.0722
	return uint8(utils.ClampUint8(math.Round(lightness / (math.MaxUint8 + 1))))
}

func ValueEval(c color.Color) uint8 {
	r, g, b, _ := c.RGBA()

	value := float64(r)*0.33 + float64(g)*0.33 + float64(b)*0.33
	return uint8(utils.ClampUint8(math.Round(value / (math.MaxUint8 + 1))))
}

func AlphaEval(c color.Color) uint8 {
	_, _, _, a := c.RGBA()

	return uint8(a >> 8)
}
