package effect

import (
	core ".."
	"../strategy"
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// SingleKernel encapsulates the logic data needed for a filter using a single kernel
// and implements the general way such a filter is applied on an image
type SingleKernel struct {
	core.BaseEffect
	strategy strategy.EdgeHandlingStrategy
	matrix   [][]float64
}

func (sk *SingleKernel) Apply(img image.Image) core.Promise {
	ret := core.CreateRGBA(img.Bounds())
	contract := sk.GetEngine().Contract(img.Bounds().Dy())
	radius := len(sk.matrix) - 1

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {

				var newRed float64
				var newGreen float64
				var newBlue float64
				for yy := -radius; yy <= radius; yy++ {
					for xx := -radius; xx <= radius; xx++ {
						r, g, b, _ := sk.strategy(&img, x + xx, y + yy).RGBA()

						newRed += float64(r) * sk.matrix[yy + radius][xx + radius]
						newGreen += float64(g) * sk.matrix[yy + radius][xx + radius]
						newBlue += float64(b) * sk.matrix[yy + radius][xx + radius]
					}
				}

				ret.(draw.Image).Set(x, y, color.RGBA64{R: uint16(core.ClampUint16(newRed)), G: uint16(core.ClampUint16(newGreen)), B: uint16(core.ClampUint16(newBlue))})
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, &contract)
}
