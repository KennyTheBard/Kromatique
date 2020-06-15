package strategy

import (
	"image/color"
	"math"
)

// ColorEvaluation is a function that transform a given color with multiple
// color channels into a single numerical value, easier to index
type ColorEvaluation func(color.Color) float64

func LightnessEval(c color.Color) float64 {
	r, g, b, _ := c.RGBA()

	lightness := float64(r)*0.2126 + float64(g)*0.7152 + float64(b)*0.0722
	return lightness / math.MaxUint16
}

func ValueEval(c color.Color) float64 {
	r, g, b, _ := c.RGBA()

	value := float64(r)*0.33 + float64(g)*0.33 + float64(b)*0.33
	return value
}

func AlphaEval(c color.Color) float64 {
	_, _, _, a := c.RGBA()

	return float64(a) / math.MaxUint16
}
