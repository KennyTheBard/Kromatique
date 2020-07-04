package effect

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/kennythebard/kromatique/core"
	"github.com/kennythebard/kromatique/strategy"
	"github.com/kennythebard/kromatique/utils"
)

// Filter encapsulates the data needed for a filter using a single kernel
// and implements the general way such a filter is applied on an image
func Convolution(edgeHandling strategy.EdgeHandling, kernel utils.Kernel) func(image.Image) image.Image {
	return func(img image.Image) image.Image {
		ret := utils.CreateRGBA(img.Bounds())
		radius := kernel.Radius()

		core.Parallelize(img.Bounds().Dy(), func(y int) {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				var newRed, newGreen, newBlue, newAlpha float64
				for yy := -radius; yy <= radius; yy++ {
					for xx := -radius; xx <= radius; xx++ {
						newX, newY := edgeHandling(img.Bounds(), x+xx, y+yy)
						r, g, b, a := img.At(newX, newY).RGBA()

						// TODO : replace this with a kernel method applied on color
						newRed += float64(r) * kernel.Get(xx+radius, yy+radius)
						newGreen += float64(g) * kernel.Get(xx+radius, yy+radius)
						newBlue += float64(b) * kernel.Get(xx+radius, yy+radius)
						newAlpha += float64(a) * kernel.Get(xx+radius, yy+radius)
					}
				}

				ret.(draw.Image).Set(x, y, color.RGBA64{
					R: uint16(utils.ClampUint16(newRed)),
					G: uint16(utils.ClampUint16(newGreen)),
					B: uint16(utils.ClampUint16(newBlue)),
					A: uint16(utils.ClampUint16(newAlpha))})
			}
		})

		return ret
	}
}

// MultiFilter encapsulates the logic data needed for a filter using multiple kernels
// and implements a customizable way of defining the behaviour
func MultiConvolution(edgeHandling strategy.EdgeHandling, merge strategy.ColorMerger, kernels ...utils.Kernel) func(image.Image) image.Image {
	return func(img image.Image) image.Image {
		ret := utils.CreateRGBA(img.Bounds())

		numKernels := len(kernels)
		radiusList := make([]int, len(kernels))
		for i, _ := range radiusList {
			radiusList[i] = kernels[i].Radius()
		}

		core.Parallelize(img.Bounds().Dy(), func(y int) {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				newColors := make([]color.Color, numKernels)

				for idx, kernel := range kernels {
					radius := radiusList[idx]
					var newRed, newGreen, newBlue, newAlpha float64
					for yy := -radius; yy <= radius; yy++ {
						for xx := -radius; xx <= radius; xx++ {
							newX, newY := edgeHandling(img.Bounds(), x+xx, y+yy)
							r, g, b, a := img.At(newX, newY).RGBA()

							newRed += float64(r) * kernel.Get(xx+radius, yy+radius)
							newGreen += float64(g) * kernel.Get(xx+radius, yy+radius)
							newBlue += float64(b) * kernel.Get(xx+radius, yy+radius)
							newAlpha += float64(a) * kernel.Get(xx+radius, yy+radius)
						}
					}

					newColors[idx] = color.RGBA64{
						R: uint16(utils.ClampUint16(math.Round(math.Abs(newRed)))),
						G: uint16(utils.ClampUint16(math.Round(math.Abs(newGreen)))),
						B: uint16(utils.ClampUint16(math.Round(math.Abs(newBlue)))),
						A: uint16(utils.ClampUint16(math.Round(math.Abs(newAlpha)))),
					}
				}

				ret.(draw.Image).Set(x, y, merge(newColors))
			}
		})

		return ret
	}
}
