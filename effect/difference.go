package effect

import (
	"image"
	"image/draw"

	"github.com/kennythebard/kromatique/core"
	"github.com/kennythebard/kromatique/strategy"
	"github.com/kennythebard/kromatique/utils"
)

// Difference returns a function that compares 2 images with the given strategy
func Difference(diff strategy.ColorDifference) func(image.Image, image.Image) image.Image {
	return func(imgA, imgB image.Image) image.Image {
		ret := utils.CreateRGBA(imgA.Bounds())

		core.Parallelize(imgA.Bounds().Dy(), func(y int) {
			for x := imgA.Bounds().Min.X; x < imgA.Bounds().Max.X; x++ {
				ret.(draw.Image).Set(x, y, diff(imgA.At(x, y), imgB.At(x, y)))
			}
		})

		return ret
	}
}
