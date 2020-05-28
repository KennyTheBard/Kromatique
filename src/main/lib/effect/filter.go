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

// EdgeHandlingStrategy defines an interface for all functions used
// to determine the behaviour of filtering around the edge of the image
type EdgeHandlingStrategy func(*image.Image, int, int) color.Color

// Extends returns the color of the closest pixel of the image
func Extend(img *image.Image, x, y int) color.Color {
	bounds := (*img).Bounds()

	if x < bounds.Min.X {
		x = bounds.Min.X
	} else if x > bounds.Max.X {
		x = bounds.Max.X
	}

	if y < bounds.Min.Y {
		y = bounds.Min.Y
	} else if y > bounds.Max.Y {
		y = bounds.Max.Y
	}

	return (*img).At(x, y)
}

// Wrap returns the color of the pixel as if the image is conceptually
// wrapped (or tiled) and values are taken from the opposite edge or corner.
func Wrap(img *image.Image, x, y int) color.Color {
	bounds := (*img).Bounds()

	if x < bounds.Min.X {
		x += bounds.Max.X - bounds.Min.X
	} else if x > bounds.Max.X {
		x -= bounds.Max.X - bounds.Min.X
	}

	if y < bounds.Min.Y {
		y += bounds.Max.Y - bounds.Min.Y
	} else if y > bounds.Max.Y {
		y -= bounds.Max.Y - bounds.Min.Y
	}

	return (*img).At(x, y)
}

// Mirror returns the color of the pixel as if the image is conceptually
// mirrored at the edges. For example, attempting to read a pixel 3 units
// outside an edge reads one 3 units inside the edge instead.
func Mirror(img *image.Image, x, y int) color.Color {
	bounds := (*img).Bounds()

	if x < bounds.Min.X {
		x = 2*bounds.Min.X - x
	} else if x > bounds.Max.X {
		x = 2*bounds.Max.X - x
	}

	if y < bounds.Min.Y {
		y = 2*bounds.Min.Y - y
	} else if y > bounds.Max.Y {
		y = 2*bounds.Max.Y - y
	}

	return (*img).At(x, y)
}

// SingleKernel encapsulates the data needed for a filter using a single kernel
// and implements the general way such a filter is applied on an image
type SingleKernel struct {
	engine       core.Engine
	edgeHandling EdgeHandlingStrategy
	kernel       utils.Matrix
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
						r, g, b, _ := effect.edgeHandling(&img, x+xx, y+yy).RGBA()

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

// ResultMergingStrategy defines an interface for all functions used
// to determine the merged result of multiple applied filters for MultiKernel
type ResultMergingStrategy func([]float64) float64

// SobelGradient defines the rule used to merge results for sobel operator;
// despite sobel using only 2 kernels, the function handles more values
func SobelGradient(results []float64) float64 {
	var sum float64
	for _, res := range results {
		sum += res * res
	}

	return math.Sqrt(sum)
}

// MultiKernel encapsulates the logic data needed for a filter using multiple kernels
// and implements a customizable way of defining the behaviour
type MultiKernel struct {
	engine        core.Engine
	edgeHandling  EdgeHandlingStrategy
	resultMerging ResultMergingStrategy
	kernels       []utils.Matrix
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
							r, g, b, _ := effect.edgeHandling(&img, x+xx, y+yy).RGBA()

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
