package blend

import (
	"fmt"
	"image"
	"image/draw"

	"../core"
	"../utils"
)

type Blender func(image.Image, image.Image, image.Point, Mask) *core.Promise

type Factory struct {
	engine core.Engine
}

func NewFactory(engine core.Engine) *Factory {
	f := new(Factory)
	f.engine = engine

	return f
}

func (f Factory) Blend(blendingStrategy BlendingStrategy) Blender {
	return func(bg, fg image.Image, fgOrigin image.Point, fgMask Mask) *core.Promise {
		ret := utils.CreateRGBA(bg.Bounds())
		contract := f.engine.Contract(bg.Bounds().Dy())

		for i := bg.Bounds().Min.Y; i < bg.Bounds().Max.Y; i++ {
			y := i
			if err := contract.PlaceOrder(func() {
				for x := bg.Bounds().Min.X; x < bg.Bounds().Max.X; x++ {
					cbg, cfg := bg.At(x, y), fg.At(x+fgOrigin.X, y+fgOrigin.Y)

					pxColor := cbg
					if image.Pt(x, y).In(fg.Bounds()) {
						pxColor = utils.PixelLERP(cbg, blendingStrategy(cbg, cfg), fgMask(x+fgOrigin.X, y+fgOrigin.Y))
					}

					ret.(draw.Image).Set(x, y, pxColor)
				}
			}); err != nil {
				fmt.Print(err)
				break
			}
		}

		return core.NewPromise(ret, contract)
	}
}

func (f Factory) Normal() Blender {
	return f.Blend(Normal)
}

func (f Factory) Difference() Blender {
	return f.Blend(Difference)
}

func (f Factory) Subtract() Blender {
	return f.Blend(Subtract)
}

func (f Factory) Darken() Blender {
	return f.Blend(Darken)
}

func (f Factory) Lighten() Blender {
	return f.Blend(Lighten)
}

func (f Factory) LinearBurn() Blender {
	return f.Blend(LinearBurn)
}

func (f Factory) Exclusion() Blender {
	return f.Blend(Exclusion)
}

func (f Factory) Divide() Blender {
	return f.Blend(Divide)
}
