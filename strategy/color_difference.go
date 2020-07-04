package strategy

import (
	"image/color"
	"math"

	"../utils"
)

// ColorDifference takes 2 colors and return a difference between them based
// on the logic it contains
type ColorDifference func(color.Color, color.Color) color.Color

// BinaryDifference returns a ColorDifference that for 2 colors that are more
// different (in matter of color channels) than a threshold, returns a predefined
// color, otherwise returns other predefined color, reducing the image to only 2 colors
func BinaryDifference(delta float64, same, difference color.Color) ColorDifference {
	border := int(math.Round(delta * math.MaxUint16))
	return func(c1, c2 color.Color) color.Color {
		r1, g1, b1, a1 := c1.RGBA()
		r2, g2, b2, a2 := c2.RGBA()
		dif := utils.Abs(int(r1)-int(r2)) + utils.Abs(int(g1)-int(g2)) +
			utils.Abs(int(b1)-int(b2)) + utils.Abs(int(a1)-int(a2))

		if dif < border {
			return same
		} else {
			return difference
		}
	}
}

// ChannelDifference returns a ColorDifference that returns the absolute value of
// the effective difference of 2 colors, channel by channel; for example, the difference
// for red and blue will be magenta, as both channels will be returned at maximum value
func ChannelDifference(c1, c2 color.Color) color.Color {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()
	return color.RGBA64{
		R: uint16(utils.Abs(int(r1) - int(r2))),
		G: uint16(utils.Abs(int(g1) - int(g2))),
		B: uint16(utils.Abs(int(b1) - int(b2))),
		A: math.MaxUint16,
	}
}
