package effect

import (
	"image"
	"image/draw"
	"math"

	"github.com/kennythebard/kromatique/core"
	"github.com/kennythebard/kromatique/strategy"
	"github.com/kennythebard/kromatique/utils"
)

// Distortion applies the Vectors obtained from the given lens
// to an image, with respect to the EdgeHandling strategy given
func Distortion(edgeHandling strategy.EdgeHandling, lens strategy.Lens) func(image.Image) image.Image {
	return func(img image.Image) image.Image {
		ret := utils.CreateRGBA(img.Bounds())

		core.Parallelize(img.Bounds().Dy(), func(y int) {
			for x := ret.Bounds().Min.X; x < ret.Bounds().Max.X; x++ {
				v := lens.VecAt(x, y)
				newX, newY := edgeHandling(img.Bounds(), int(math.Round(float64(x)+v.X)), int(math.Round(float64(y)+v.Y)))

				ret.(draw.Image).Set(x, y, img.At(newX, newY))
			}
		})

		return ret
	}
}
