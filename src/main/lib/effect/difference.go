package effect

import (
	"fmt"
	"image"
	"image/draw"

	"../core"
	"../strategy"
	"../utils"
)

// Difference serves as a generic customizable structure that encapsulates
// the logic needed to apply a given difference strategy
type Difference struct {
	engine core.Engine
	diff   strategy.Difference
}

func (effect *Difference) Apply(imgA, imgB image.Image) *core.Promise {
	ret := utils.CreateRGBA(imgA.Bounds())
	contract := effect.engine.Contract(imgA.Bounds().Dy())

	for i := imgA.Bounds().Min.Y; i < imgA.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := imgA.Bounds().Min.X; x < imgA.Bounds().Max.X; x++ {
				ret.(draw.Image).Set(x, y, effect.diff(imgA.At(x, y), imgB.At(x, y)))
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}
