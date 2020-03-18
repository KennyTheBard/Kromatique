package mirror

import (
	"fmt"
	"image"
	"image/draw"

	core "../.."
	"../../utils"
)

type Reflect func(int, int, image.Rectangle) (int, int)

func HorizontalReflect(x, y int, bounds image.Rectangle) (int, int) {
	return bounds.Max.X - (x - bounds.Min.X), y
}

func VerticalReflect(x, y int, bounds image.Rectangle) (int, int) {
	return x, bounds.Max.Y - (y - bounds.Min.Y)
}

type Mirror struct {
	core.Base
	reflect Reflect
}

func (effect *Mirror) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				newX, newY := effect.reflect(x, y, img.Bounds())
				ret.(draw.Image).Set(x, y, img.At(newX, newY))
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}

func NewMirror(reflect Reflect) *Mirror {
	m := new(Mirror)
	m.reflect = reflect

	return m
}

func NewHorizontalMirror() *Mirror {
	return NewMirror(HorizontalReflect)
}

func NewVerticalMirror() *Mirror {
	return NewMirror(VerticalReflect)
}
