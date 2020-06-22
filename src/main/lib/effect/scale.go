package effect

import (
	"../core"
	"../strategy"
	"../utils"
	"image"
	"image/draw"
)

// Scale serves as a generic customizable structure that encapsulates
// the logic needed to apply a scaling transformation on an image
type Scale struct {
	engine                core.Engine
	colorSamplingStrategy strategy.ColorSampling
}

func (effect *Scale) Apply(img image.Image, dx, dy int) *core.Promise {
	factor := strategy.ScaleFactor{
		X: float64(dx) / float64(img.Bounds().Dx()),
		Y: float64(dy) / float64(img.Bounds().Dy()),
	}
	trgBounds := image.Rect(0, 0, dx, dy)

	ret := utils.CreateRGBA(trgBounds)
	contract := effect.engine.Contract()

	for i := ret.Bounds().Min.Y; i < ret.Bounds().Max.Y; i++ {
		y := i
		contract.PlaceOrder(func() {
			for x := ret.Bounds().Min.X; x < ret.Bounds().Max.X; x++ {
				col := effect.colorSamplingStrategy(img, x, y, factor)

				ret.(draw.Image).Set(x, y, col)
			}
		})
	}

	return contract.Promise(ret)
}
