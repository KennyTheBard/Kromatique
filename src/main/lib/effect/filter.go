package effect

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"../core"
	"../strategy"
	"../utils"
)

// SingleKernel encapsulates the data needed for a filter using a single kernel
// and implements the general way such a filter is applied on an image
type SingleKernel struct {
	engine       core.Engine
	edgeHandling strategy.EdgeHandling
	kernel       utils.Kernel
}

func (effect *SingleKernel) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.engine.Contract(img.Bounds().Dy())
	radius := effect.kernel.Radius()

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				var newRed float64
				var newGreen float64
				var newBlue float64
				for yy := -radius; yy <= radius; yy++ {
					for xx := -radius; xx <= radius; xx++ {
						newX, newY := effect.edgeHandling(img.Bounds(), x+xx, y+yy)
						r, g, b, _ := img.At(newX, newY).RGBA()

						newRed += float64(r) * effect.kernel.Get(xx+radius, yy+radius)
						newGreen += float64(g) * effect.kernel.Get(xx+radius, yy+radius)
						newBlue += float64(b) * effect.kernel.Get(xx+radius, yy+radius)
					}
				}

				ret.(draw.Image).Set(x, y, color.RGBA64{
					R: uint16(utils.ClampUint16(newRed)),
					G: uint16(utils.ClampUint16(newGreen)),
					B: uint16(utils.ClampUint16(newBlue))})
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}

// MultiKernel encapsulates the logic data needed for a filter using multiple kernels
// and implements a customizable way of defining the behaviour
type MultiKernel struct {
	engine        core.Engine
	edgeHandling  strategy.EdgeHandling
	resultMerging strategy.ResultMerger
	kernels       []utils.Kernel
}

func (effect *MultiKernel) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.engine.Contract(img.Bounds().Dy())

	radiusNumber := len(effect.kernels)
	radiusList := make([]int, len(effect.kernels))
	for i, _ := range radiusList {
		radiusList[i] = effect.kernels[i].Radius()
	}

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				newRed := make([]float64, radiusNumber)
				newGreen := make([]float64, radiusNumber)
				newBlue := make([]float64, radiusNumber)

				for index, kernel := range effect.kernels {
					radius := radiusList[index]
					for yy := -radius; yy <= radius; yy++ {
						for xx := -radius; xx <= radius; xx++ {
							newX, newY := effect.edgeHandling(img.Bounds(), x+xx, y+yy)
							r, g, b, _ := img.At(newX, newY).RGBA()

							newRed[index] += float64(r) * kernel.Get(xx+radius, yy+radius)
							newGreen[index] += float64(g) * kernel.Get(xx+radius, yy+radius)
							newBlue[index] += float64(b) * kernel.Get(xx+radius, yy+radius)
						}
					}
				}

				ret.(draw.Image).Set(x, y,
					color.RGBA64{
						R: uint16(utils.ClampUint16(effect.resultMerging(newRed))),
						G: uint16(utils.ClampUint16(effect.resultMerging(newGreen))),
						B: uint16(utils.ClampUint16(effect.resultMerging(newBlue))),
					})
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}
