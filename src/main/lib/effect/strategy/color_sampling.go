package strategy

import (
	utils "../../utils"
	"image"
	"image/color"
	"math"
)

type ColorSamplingStrategy func(int, int, ScaleFactor, image.Image) color.Color

func SinglePixelSampling(x, y int, factor ScaleFactor, img image.Image) color.Color {
	return img.At(int(math.Round(float64(x)/factor.X)), int(math.Round(float64(y)/factor.Y)))
}

func CornerPixelsSampling(x, y int, factor ScaleFactor, img image.Image) color.Color {
	exactX := float64(x) / factor.X
	exactY := float64(y) / factor.Y

	return utils.PixelBiLERP(
		img.At(int(math.Floor(exactX)), int(math.Floor(exactY))),
		img.At(int(math.Ceil(exactX)), int(math.Floor(exactY))),
		img.At(int(math.Floor(exactX)), int(math.Ceil(exactY))),
		img.At(int(math.Ceil(exactX)), int(math.Ceil(exactY))),
		exactX-math.Floor(exactX),
		exactY-math.Floor(exactY))
}
