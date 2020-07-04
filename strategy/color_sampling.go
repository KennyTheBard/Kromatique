package strategy

import (
	"image"
	"image/color"
	"math"

	"github.com/kennythebard/kromatique/utils"
)

// ColorSampling encapsulates a strategy used to decide what color
// the pixel at the given position from the destination image will have
type ColorSampling func(image.Image, int, int, ScaleFactor) color.Color

// SinglePixelSampling is an implementation of ColorSampling that uses
// the color of a single pixel (the closest to the target) from the source image
func SinglePixelSampling(img image.Image, destX, destY int, factor ScaleFactor) color.Color {
	exactX, exactY := factor.GetSourcePx(float64(destX), float64(destY))

	return img.At(int(math.Round(exactX)), int(math.Round(exactY)))
}

// CornerPixelsSampling is an implementation of ColorSampling that uses
// a bilinear interpolation between the corners of the rectangle obtained by scaling
// back a pixel from the destination image to the source image; the scaled rectangle
// can actually be smaller than a single pixel would for factors greater than 1
func CornerPixelsSampling(img image.Image, x, y int, factor ScaleFactor) color.Color {
	exactX, exactY := factor.GetSourcePx(float64(x), float64(y))

	return utils.PixelBiLERP(
		img.At(int(math.Floor(exactX)), int(math.Floor(exactY))),
		img.At(int(math.Ceil(exactX)), int(math.Floor(exactY))),
		img.At(int(math.Floor(exactX)), int(math.Ceil(exactY))),
		img.At(int(math.Ceil(exactX)), int(math.Ceil(exactY))),
		exactX-math.Floor(exactX),
		exactY-math.Floor(exactY))
}
