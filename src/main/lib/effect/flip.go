package effect

import (
	"image"
	"image/draw"

	"../core"
	"../utils"
)

// FlipperStrategy returns the flipped position for a given position in the bounds of the image
type FlipperStrategy func(int, int, image.Rectangle) (int, int)

// HorizontalFlip returns the given position flipped horizontally
func HorizontalFlip(x, y int, bounds image.Rectangle) (int, int) {
	return bounds.Max.X - (x - bounds.Min.X), y
}

// VerticalFlip returns the given position flipped vertically
func VerticalFlip(x, y int, bounds image.Rectangle) (int, int) {
	return x, bounds.Max.Y - (y - bounds.Min.Y)
}

// Flip serves as a generic customizable structure that encapsulates
// the logic needed to apply a flipping strategy
type Flip struct {
	engine   core.Engine
	strategy FlipperStrategy
}

func (effect *Flip) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.engine.Contract()

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				newX, newY := effect.strategy(x, y, img.Bounds())
				ret.(draw.Image).Set(x, y, img.At(newX, newY))
			}
		})
	}

	return contract.Promise(ret)
}
