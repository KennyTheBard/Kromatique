package effect

import (
	"image"
	"image/color"
	"image/draw"
	"sort"

	"../core"
	"../strategy"
	"../utils"
)

// Median serves as a generic customizable structure that encapsulates
// the logic needed to apply a a median effect on an image
type Median struct {
	engine       core.Engine
	eval         strategy.ColorEvaluation
	edgeHandling strategy.EdgeHandling
	windowRadius int
}

func (effect *Median) Apply(img image.Image) *core.Promise {
	bounds := img.Bounds()
	ret := utils.CreateRGBA(bounds)
	contract := effect.engine.Contract()

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				radius := effect.windowRadius
				windowValues := make([]color.Color, (2*radius+1)*(2*radius+1))
				for a := -radius; a <= radius; a++ {
					for b := -radius; b <= radius; b++ {
						newX, newY := effect.edgeHandling(bounds, x+a, y+b)
						windowValues[(a+radius)*(2*radius+1)+(b+radius)] = img.At(newX, newY)
					}
				}

				sort.Slice(windowValues[:], func(i, j int) bool {
					return effect.
						eval(windowValues[i]) < effect.eval(windowValues[j])
				})

				ret.(draw.Image).Set(x, y, windowValues[(2*radius+1)*(2*radius+1)/2])
			}
		})
	}

	return contract.Promise(ret)
}
