package mapper

import (
	core "../.."
	"../../utils"
	"fmt"
	"image"
	"image/draw"
)

type ColorMapperRunner struct {
	core.Base
	mappers []ColorMapper
}

func (effect *ColorMapperRunner) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				currentColor := img.At(x, y)
				for _, mapper := range effect.mappers {
					currentColor = mapper(currentColor)
				}

				ret.(draw.Image).Set(x, y, currentColor)
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}

func (effect *ColorMapperRunner) Add(mapper ColorMapper) {
	effect.mappers = append(effect.mappers, mapper)
}

func NewColorMapperRunner() *ColorMapperRunner {
	cmr := new(ColorMapperRunner)
	cmr.mappers = make([]ColorMapper, 0)

	return cmr
}
