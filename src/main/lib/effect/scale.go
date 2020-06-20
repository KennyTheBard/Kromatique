package effect

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"../core"
	"../strategy"
	"../utils"
)

// ColorSamplingStrategy encapsulates a strategy used to decide what color
// the pixel at the given position from the destination image will have
type ColorSamplingStrategy func(image.Image, int, int, strategy.ScaleFactor) color.Color

// SinglePixelSampling is an implementation of ColorSamplingStrategy that uses
// the color of a single pixel (the closest to the target) from the source image
func SinglePixelSampling(img image.Image, destX, destY int, factor strategy.ScaleFactor) color.Color {
	exactX, exactY := factor.GetSourcePx(float64(destX), float64(destY))

	return img.At(int(math.Round(exactX)), int(math.Round(exactY)))
}

// CornerPixelsSampling is an implementation of ColorSamplingStrategy that uses
// a bilinear interpolation between the corners of the rectangle obtained by scaling
// back a pixel from the destination image to the source image; the scaled rectangle
// can actually be smaller than a single pixel would for factors greater than 1
func CornerPixelsSampling(img image.Image, x, y int, factor strategy.ScaleFactor) color.Color {
	exactX, exactY := factor.GetSourcePx(float64(x), float64(y))

	return utils.PixelBiLERP(
		img.At(int(math.Floor(exactX)), int(math.Floor(exactY))),
		img.At(int(math.Ceil(exactX)), int(math.Floor(exactY))),
		img.At(int(math.Floor(exactX)), int(math.Ceil(exactY))),
		img.At(int(math.Ceil(exactX)), int(math.Ceil(exactY))),
		exactX-math.Floor(exactX),
		exactY-math.Floor(exactY))
}

// Scale serves as a generic customizable structure that encapsulates
// the logic needed to apply a scaling transformation on an image
type Scale struct {
	engine                core.Engine
	scaleFactorStrategy   strategy.ScalingTarget
	colorSamplingStrategy ColorSamplingStrategy
}

func (effect *Scale) Apply(img image.Image) *core.Promise {
	factor, trgBounds := effect.scaleFactorStrategy(img.Bounds())

	ret := utils.CreateRGBA(trgBounds)
	contract := effect.engine.Contract()

	for i := ret.Bounds().Min.Y; i < ret.Bounds().Max.Y; i++ {
		y := i
		contract.PlaceOrder(func() {
			for x := ret.Bounds().Min.X; x < ret.Bounds().Max.X; x++ {
				col := effect.colorSamplingStrategy(img, x, y, factor)

				ret.(draw.Image).Set(x, y, col)
			}
		})
	}

	return contract.Promise(ret)
}
