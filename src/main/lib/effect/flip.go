package effect

import (
	"fmt"
	"image"
	"image/draw"

	core ".."
	"../utils"
)

// FlipStrategy returns the flipped position for a given position in the bounds of the image
type FlipStrategy func(int, int, image.Rectangle) (int, int)

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
	core.Base
	strategy FlipStrategy
}

func (effect *Flip) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				newX, newY := effect.strategy(x, y, img.Bounds())
				ret.(draw.Image).Set(x, y, img.At(newX, newY))
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}

func NewFlip(strategy FlipStrategy) *Flip {
	m := new(Flip)
	m.strategy = strategy

	return m
}

func NewHorizontalFlip() *Flip {
	return NewFlip(HorizontalFlip)
}

func NewVerticalFlip() *Flip {
	return NewFlip(VerticalFlip)
}
