package overlay

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	core "../.."
	"../../position"
	"../../utils"
)

// Overlay encapsulates the data needed to apply an overlay image
type Overlay struct {
	core.Base
	Stamp   image.Image
	Origin  position.Position
	Opacity float64
}

func (effect *Overlay) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())
	stampBounds := effect.Stamp.Bounds()
	origin := effect.Origin.Get(img.Bounds())

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				if x < origin.X || y < origin.Y ||
					x > stampBounds.Max.X+origin.X ||
					y > stampBounds.Max.Y+origin.Y {
					ret.(draw.Image).Set(x, y, img.At(x, y))
					continue
				}

				or, og, ob, oa := effect.Stamp.At(x-origin.X, y-origin.Y).RGBA()

				if oa == 0 {
					ret.(draw.Image).Set(x, y, img.At(x, y))
				} else {
					opacity := utils.LERP(float64(oa)/utils.MaxUint16, 0.0, effect.Opacity)

					if opacity == utils.MaxUint16 {
						ret.(draw.Image).Set(x, y, color.RGBA64{R: uint16(or), G: uint16(og), B: uint16(ob)})
					} else {
						r, g, b, a := img.At(x, y).RGBA()

						newRed := utils.LERP(float64(or), float64(r), opacity)
						newGreen := utils.LERP(float64(og), float64(g), opacity)
						newBlue := utils.LERP(float64(ob), float64(b), opacity)

						ret.(draw.Image).Set(x, y, color.RGBA64{
							R: uint16(utils.ClampUint16(newRed)),
							G: uint16(utils.ClampUint16(newGreen)),
							B: uint16(utils.ClampUint16(newBlue)),
							A: uint16(a),
						})
					}
				}
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}
