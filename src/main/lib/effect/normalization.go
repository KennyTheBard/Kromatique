package effect

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"

	"../core"
	"../utils"
)

// Normalization serves as a generic customizable structure that encapsulates
// the logic needed to normalize the color values of an image
type Normalization struct {
	engine                         core.Engine
	sourceInterval, targetInterval *utils.ColorInterval
}

func (effect *Normalization) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.engine.Contract(img.Bounds().Dy())
	ratio := float64(effect.targetInterval.Max()-effect.targetInterval.Min()) /
		float64(effect.sourceInterval.Max()-effect.sourceInterval.Min())

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				I := (r + g + b) / 3
				newI := float64(I-effect.sourceInterval.Min())*ratio +
					float64(effect.targetInterval.Min())
				ret.(draw.Image).Set(x, y, color.Gray16{Y: uint16(utils.ClampUint16(math.Round(newI)))})
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}
