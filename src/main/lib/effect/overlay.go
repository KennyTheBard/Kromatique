package effect

import (
	core ".."
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// Overlay encapsulates the data needed to apply an overlay image
type Overlay struct {
	core.BaseEffect
	Stamp        image.Image
	Origin       image.Point
	//transparency float64
}

func (effect *Overlay) Apply(img image.Image) core.Promise {
	ret := core.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())
	stampBounds := effect.Stamp.Bounds()

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				if x < effect.Origin.X ||
					y < effect.Origin.Y ||
					x > stampBounds.Max.X+effect.Origin.X ||
					y > stampBounds.Max.Y+effect.Origin.Y {
					ret.(draw.Image).Set(x, y, img.At(x, y))
					continue
				}

				or, og, ob, oa := effect.Stamp.At(x-effect.Origin.X, y-effect.Origin.Y).RGBA()

				if oa == 0 {
					ret.(draw.Image).Set(x, y, img.At(x, y))
				} else if oa == core.MaxUint16 {
					ret.(draw.Image).Set(x, y, color.RGBA64{R: uint16(or), G: uint16(og), B: uint16(ob)})
				} else {
					r, g, b, a := img.At(x, y).RGBA()

					newRed := core.Lerp(float64(or), float64(r), float64(oa) / core.MaxUint16)
					newGreen := core.Lerp(float64(og), float64(g), float64(oa) / core.MaxUint16)
					newBlue := core.Lerp(float64(ob), float64(b), float64(oa) / core.MaxUint16)

					ret.(draw.Image).Set(x, y, color.RGBA64{
						R: uint16(core.ClampUint16(newRed)),
						G: uint16(core.ClampUint16(newGreen)),
						B: uint16(core.ClampUint16(newBlue)),
						A: uint16(a),
					})
				}
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, &contract)
}
