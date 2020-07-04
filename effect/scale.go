package effect

import (
	"../core"
	"../strategy"
	"../utils"
	"image"
	"image/draw"
)

// Scale serves as a generic customizable structure that encapsulates
// the logic needed to apply a scaling transformation on an image
func Scale(colorSamplingStrategy strategy.ColorSampling) func(image.Image, int, int) image.Image {
	return func(img image.Image, dx, dy int) image.Image {
		factor := strategy.ScaleFactor{
			X: float64(dx) / float64(img.Bounds().Dx()),
			Y: float64(dy) / float64(img.Bounds().Dy()),
		}
		trgBounds := image.Rect(0, 0, dx, dy)

		ret := utils.CreateRGBA(trgBounds)

		core.Parallelize(trgBounds.Dy(), func(y int) {
			for x := ret.Bounds().Min.X; x < ret.Bounds().Max.X; x++ {
				col := colorSamplingStrategy(img, x, y, factor)

				ret.(draw.Image).Set(x, y, col)
			}
		})

		return ret
	}
}
