package scale

import (
	core "../.."
	utils "../../utils"
	"fmt"
	"image"
	"image/draw"
)

type Scale struct {
	core.Base
	scaleFactorStrategy   ScaleFactorStrategy
	colorSamplingStrategy ColorSamplingStrategy
}

func (effect *Scale) Apply(img image.Image) *core.Promise {
	factor := effect.scaleFactorStrategy.Factor(img.Bounds())
	trgBounds := effect.scaleFactorStrategy.Size(img.Bounds())

	ret := utils.CreateRGBA(trgBounds)
	contract := effect.GetEngine().Contract(ret.Bounds().Dy())

	for i := ret.Bounds().Min.Y; i < ret.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := ret.Bounds().Min.X; x < ret.Bounds().Max.X; x++ {
				col := effect.colorSamplingStrategy(img, x, y, factor)

				ret.(draw.Image).Set(x, y, col)
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}

func NewScale(scaleFactorStrategy ScaleFactorStrategy, colorSamplingStrategy ColorSamplingStrategy) *Scale {
	scale := new(Scale)
	scale.scaleFactorStrategy = scaleFactorStrategy
	scale.colorSamplingStrategy = colorSamplingStrategy

	return scale
}
