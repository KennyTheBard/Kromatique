package effect

import (
	"image"
	"image/draw"

	"github.com/kennythebard/kromatique/core"
	"github.com/kennythebard/kromatique/strategy"
	"github.com/kennythebard/kromatique/utils"
)

// Scale returns a function that uses the given ColorSampling strategy
// in order to resize an image to the required dimensions
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
