package effect

import (
	wrap "./wrap"
	core ".."
	"./strategy"
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// SingleKernel encapsulates the data needed for a filter using a single kernel
// and implements the general way such a filter is applied on an image
type SingleKernel struct {
	core.BaseEffect
	EdgeHandling strategy.EdgeHandlingStrategy
	Kernel       wrap.Matrix
}

func (effect *SingleKernel) Apply(img image.Image) core.Promise {
	ret := core.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())
	radius := effect.Kernel.Radius()

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				var newRed float64
				var newGreen float64
				var newBlue float64
				for yy := -radius; yy <= radius; yy++ {
					for xx := -radius; xx <= radius; xx++ {
						r, g, b, _ := effect.EdgeHandling(&img, x+xx, y+yy).RGBA()

						newRed += float64(r) * effect.Kernel.Get(xx+radius, yy+radius)
						newGreen += float64(g) * effect.Kernel.Get(xx+radius, yy+radius)
						newBlue += float64(b) * effect.Kernel.Get(xx+radius, yy+radius)
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

// MultiKernel encapsulates the logic data needed for a filter using multiple kernels
// and implements a customizable way of defining the behaviour
type MultiKernel struct {
	core.BaseEffect
	EdgeHandling  strategy.EdgeHandlingStrategy
	ResultMerging strategy.ResultMergingStrategy
	Kernels       []wrap.Matrix
}

func (effect *MultiKernel) Apply(img image.Image) core.Promise {
	ret := core.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())

	radiusNumber := len(effect.Kernels)
	radiusList := make([]int, len(effect.Kernels))
	for i, _ := range radiusList {
		radiusList[i] = effect.Kernels[i].Radius()
	}

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				newRed := make([]float64, radiusNumber)
				newGreen := make([]float64, radiusNumber)
				newBlue := make([]float64, radiusNumber)

				for index, kernel := range effect.Kernels {
					radius := radiusList[index]
					for yy := -radius; yy <= radius; yy++ {
						for xx := -radius; xx <= radius; xx++ {
							r, g, b, _ := effect.EdgeHandling(&img, x+xx, y+yy).RGBA()

							newRed[index] += float64(r) * kernel.Get(xx+radius, yy+radius)
							newGreen[index] += float64(g) * kernel.Get(xx+radius, yy+radius)
							newBlue[index] += float64(b) * kernel.Get(xx+radius, yy+radius)
						}
					}
				}

				ret.(draw.Image).Set(x, y,
					color.RGBA64{
						R: uint16(core.ClampUint16(effect.ResultMerging(newRed))),
						G: uint16(core.ClampUint16(effect.ResultMerging(newGreen))),
						B: uint16(core.ClampUint16(effect.ResultMerging(newBlue))),
					})
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, &contract)
}
