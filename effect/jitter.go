package effect

import (
	"image"
	"image/draw"
	"math/rand"
	"time"

	"github.com/kennythebard/kromatique/core"
	"github.com/kennythebard/kromatique/utils"
)

// Jitter serves as a generic customizable structure that encapsulates
// the logic needed to apply a jitter effect on an image
func Jitter(radius int) func(image.Image) image.Image {
	return func(img image.Image) image.Image {
		ret := utils.CreateRGBA(img.Bounds())
		rand.Seed(time.Now().Unix())
		randCoordinate := func() int {
			return rand.Intn(radius*2) - radius
		}

		core.Parallelize(img.Bounds().Dy(), func(y int) {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				offsetX, offsetY := randCoordinate(), randCoordinate()
				newX := utils.Max(utils.Min(x+offsetX, img.Bounds().Max.X), img.Bounds().Min.X)
				newY := utils.Max(utils.Min(y+offsetY, img.Bounds().Max.Y), img.Bounds().Min.Y)

				ret.(draw.Image).Set(x, y, img.At(newX, newY))
			}
		})

		return ret
	}
}
