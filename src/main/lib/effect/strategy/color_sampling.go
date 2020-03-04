package strategy

import (
	utils "../../utils"
	"image"
	"image/color"
	"math"
)

type ColorSamplingStrategy func(int, int, ScaleFactor, image.Image) color.Color

func SinglePixelSampling(destX, destY int, factor ScaleFactor, img image.Image) color.Color {
	exactX, exactY := factor.ToSource(float64(destX), float64(destY))

	return img.At(int(math.Round(exactX)), int(math.Round(exactY)))
}

func CornerPixelsSampling(x, y int, factor ScaleFactor, img image.Image) color.Color {
	exactX, exactY := factor.ToSource(float64(x), float64(y))

	return utils.PixelBiLERP(
		img.At(int(math.Floor(exactX)), int(math.Floor(exactY))),
		img.At(int(math.Ceil(exactX)), int(math.Floor(exactY))),
		img.At(int(math.Floor(exactX)), int(math.Ceil(exactY))),
		img.At(int(math.Ceil(exactX)), int(math.Ceil(exactY))),
		exactX-math.Floor(exactX),
		exactY-math.Floor(exactY))
}
