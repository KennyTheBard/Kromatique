package effect

import (
	"image"
	"image/color"
	"image/draw"
	"sort"

	"github.com/kennythebard/kromatique/core"
	"github.com/kennythebard/kromatique/strategy"
	"github.com/kennythebard/kromatique/utils"
)

// Median returns a function that applies a Median effects with respect
// to the given ColorEvaluation strategy, EdgeHandling strategy and window radius to an image
func Median(eval strategy.ColorEvaluation, edgeHandling strategy.EdgeHandling, windowRadius int) func(image.Image) image.Image {
	return func(img image.Image) image.Image {
		bounds := img.Bounds()
		ret := utils.CreateRGBA(bounds)

		core.Parallelize(bounds.Dy(), func(y int) {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				radius := windowRadius
				windowValues := make([]color.Color, (2*radius+1)*(2*radius+1))
				for a := -radius; a <= radius; a++ {
					for b := -radius; b <= radius; b++ {
						newX, newY := edgeHandling(bounds, x+a, y+b)
						windowValues[(a+radius)*(2*radius+1)+(b+radius)] = img.At(newX, newY)
					}
				}

				sort.Slice(windowValues[:], func(i, j int) bool {
					return eval(windowValues[i]) < eval(windowValues[j])
				})

				ret.(draw.Image).Set(x, y, windowValues[(2*radius+1)*(2*radius+1)/2])
			}
		})

		return ret
	}
}
