package effect

import (
	"image"
	"image/draw"
	"math"

	"../core"
	"../strategy"
	"../utils"
)

// Distortion serves as a generic customizable structure that encapsulates
// the logic needed to apply a distortion on a given image
type Distortion struct {
	engine       core.Engine
	edgeHandling strategy.EdgeHandling
	lens         strategy.Lens
}

func (effect *Distortion) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.engine.Contract()

	for i := ret.Bounds().Min.Y; i < ret.Bounds().Max.Y; i++ {
		y := i
		contract.PlaceOrder(func() {
			for x := ret.Bounds().Min.X; x < ret.Bounds().Max.X; x++ {
				v := effect.lens.VecAt(x, y)
				newX, newY := effect.edgeHandling(img.Bounds(), int(math.Round(float64(x)+v.X)), int(math.Round(float64(y)+v.Y)))

				ret.(draw.Image).Set(x, y, img.At(newX, newY))
			}
		})
	}

	return contract.Promise(ret)
}
