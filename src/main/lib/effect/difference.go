package effect

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"

	"../core"
	"../utils"
)

// DifferenceStrategy takes 2 colors and return a difference between them based
// on the logic it contains
type DifferenceStrategy func(color.Color, color.Color) color.Color

// BinaryDifference returns a DifferenceStrategy that for 2 colors that are more
// different (in matter of color channels) than a threshold, returns a predefined
// color, otherwise returns other predefined color, reducing the image to only 2 colors
func BinaryDifference(delta float64, same, difference color.Color) DifferenceStrategy {
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

// ChannelDifference returns a DifferenceStrategy that returns the absolute value of
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

// Difference serves as a generic customizable structure that encapsulates
// the logic needed to apply a given difference strategy
type Difference struct {
	engine core.Engine
	diff   DifferenceStrategy
}

func (effect *Difference) Apply(imgA, imgB image.Image) *core.Promise {
	ret := utils.CreateRGBA(imgA.Bounds())
	contract := effect.engine.Contract(imgA.Bounds().Dy())

	for i := imgA.Bounds().Min.Y; i < imgA.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := imgA.Bounds().Min.X; x < imgA.Bounds().Max.X; x++ {
				ret.(draw.Image).Set(x, y, effect.diff(imgA.At(x, y), imgB.At(x, y)))
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}
