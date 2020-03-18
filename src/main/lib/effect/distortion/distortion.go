package distorsion

import (
	"fmt"
	"image"
	"image/draw"
	"math"

	core "../.."
	"../../utils"
	"../filter"
)

type Distortion struct {
	core.Base
	edgeHandling filter.EdgeHandlingStrategy
	asm          LensAssembly
}

func (effect *Distortion) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(ret.Bounds().Dy())
	vm := effect.asm.VectorMap()

	for i := ret.Bounds().Min.Y; i < ret.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := ret.Bounds().Min.X; x < ret.Bounds().Max.X; x++ {
				if image.Pt(x, y).In(vm.Bounds()) {
					v := vm.At(x, y)
					newX := int(math.Round(float64(x) + v.X))
					newY := int(math.Round(float64(y) + v.Y))
					col := effect.edgeHandling(&img, newX, newY)

					ret.(draw.Image).Set(x, y, col)
				} else {
					ret.(draw.Image).Set(x, y, img.At(x, y))
				}
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}

func NewDistortion(edgeHandling filter.EdgeHandlingStrategy, asm LensAssembly) *Distortion {
	d := new(Distortion)
	d.edgeHandling = edgeHandling
	d.asm = asm

	return d
}
