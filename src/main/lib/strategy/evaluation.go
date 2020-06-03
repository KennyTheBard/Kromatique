package strategy

import (
	"image/color"
	"math"
)

// Evaluation is a function that transform a given color with multiple
// color channels into a single numerical value, easier to index
type Evaluation func(color.Color) uint

func LightnessEval(c color.Color) uint {
	r, g, b, _ := c.RGBA()

	lightness := float64(r>>8)*0.2126 + float64(g>>8)*0.7152 + float64(b>>8)*0.0722
	return uint(math.Round(lightness))
}

func RedEval(c color.Color) uint {
	r, _, _, _ := c.RGBA()

	return uint(r >> 8)
}

func GreenEval(c color.Color) uint {
	_, g, _, _ := c.RGBA()

	return uint(g >> 8)
}

func BlueEval(c color.Color) uint {
	_, _, b, _ := c.RGBA()

	return uint(b >> 8)
}

func AlphaEval(c color.Color) uint {
	_, _, _, a := c.RGBA()

	return uint(a >> 8)
}
