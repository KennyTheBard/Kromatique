package effect

import (
	"image"
	"image/draw"

	"../core"
	"../strategy"
	"../utils"
)

// ColorMapper serves as a generic customizable structure that encapsulates
// the logic needed to apply a series of MappingRule on an image
type ColorMapper struct {
	engine core.Engine
	rules  []strategy.MappingRule
}

func (effect *ColorMapper) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.engine.Contract()

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				currentColor := img.At(x, y)
				for _, rule := range effect.rules {
					currentColor = rule(currentColor)
				}

				ret.(draw.Image).Set(x, y, currentColor)
			}
		})
	}

	return contract.Promise(ret)
}
