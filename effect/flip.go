package effect

import (
	"image"
	"image/draw"

	"github.com/kennythebard/kromatique/core"
	"github.com/kennythebard/kromatique/utils"
)

// FlipperStrategy returns the flipped position for a given position in the bounds of the image
type FlipperStrategy func(int, int, image.Rectangle) (int, int)

// HorizontalFlip returns the given position flipped horizontally
func HorizontalFlipper(x, y int, bounds image.Rectangle) (int, int) {
	return bounds.Max.X - (x - bounds.Min.X), y
}

// VerticalFlip returns the given position flipped vertically
func VerticalFlipper(x, y int, bounds image.Rectangle) (int, int) {
	return x, bounds.Max.Y - (y - bounds.Min.Y)
}

// Flip returns the given image.Image flipper by given FlipperStrategy
func Flip(img image.Image, strategy FlipperStrategy) image.Image {
	ret := utils.CreateRGBA(img.Bounds())

	core.Parallelize(img.Bounds().Dy(), func(y int) {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			newX, newY := strategy(x, y, img.Bounds())
			ret.(draw.Image).Set(x, y, img.At(newX, newY))
		}
	})

	return ret
}

// FlipHorizontal returns the given image.Image flipper horizontally
func FlipHorizontal(img image.Image) image.Image {
	return Flip(img, HorizontalFlipper)
}

// FlipHorizontal returns the given image.Image flipper vertically
func FlipVertical(img image.Image) image.Image {
	return Flip(img, VerticalFlipper)
}
