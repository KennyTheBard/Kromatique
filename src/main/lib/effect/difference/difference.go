package difference

import (
	core "../.."
	"../../utils"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
)

type DifferenceStrategy func(color.Color, color.Color) color.Color

func BinaryDifferenceFactory(delta float64, same, difference color.Color) DifferenceStrategy {
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

func ColorChannelDifferenceStrategy(c1, c2 color.Color) color.Color {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()
	return color.RGBA64{
		R: uint16(utils.Abs(int(r1) - int(r2))),
		G: uint16(utils.Abs(int(g1) - int(g2))),
		B: uint16(utils.Abs(int(b1) - int(b2))),
		A: math.MaxUint16,
	}

}

type Difference struct {
	core.Base
	diff DifferenceStrategy
}

func (effect *Difference) Apply(imgA, imgB image.Image) *core.Promise {
	ret := utils.CreateRGBA(imgA.Bounds())
	contract := effect.GetEngine().Contract(imgA.Bounds().Dy())

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

func NewDifference(diff DifferenceStrategy) *Difference {
	d := new(Difference)
	d.diff = diff

	return d
}
