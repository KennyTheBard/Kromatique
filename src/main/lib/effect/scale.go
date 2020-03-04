package effect

import (
	core ".."
	utils "../utils"
	strategy "./strategy"
	"fmt"
	"image"
	"image/draw"
)

type Scale struct {
	core.BaseEffect
	scaleFactorStrategy   strategy.ScaleFactorStrategy
	colorSamplingStrategy strategy.ColorSamplingStrategy
}

func (effect *Scale) Apply(img image.Image) core.Promise {
	factor := effect.scaleFactorStrategy.Factor(img.Bounds())
	trgBounds := effect.scaleFactorStrategy.Size(img.Bounds())

	ret := utils.CreateRGBA(trgBounds)
	contract := effect.GetEngine().Contract(ret.Bounds().Dy())

	for i := ret.Bounds().Min.Y; i < ret.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := ret.Bounds().Min.X; x < ret.Bounds().Max.X; x++ {
				col := effect.colorSamplingStrategy(x, y, factor, img)

				ret.(draw.Image).Set(x, y, col)
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, &contract)
}

func NewScale(scaleFactorStrategy strategy.ScaleFactorStrategy, colorSamplingStrategy strategy.ColorSamplingStrategy) Scale {
	return Scale{scaleFactorStrategy: scaleFactorStrategy, colorSamplingStrategy: colorSamplingStrategy}
}
