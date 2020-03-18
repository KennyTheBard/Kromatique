package overlay

import (
	core "../.."
	"../../position"
	"../../utils"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Overlay encapsulates the data needed to apply an overlay image
type Overlay struct {
	core.Base
	stamp   image.Image
	origin  position.Position
	opacity float64
}

func (effect *Overlay) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())
	stampBounds := effect.stamp.Bounds()
	origin := effect.origin.Get(img.Bounds())

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

				or, og, ob, oa := effect.stamp.At(x-origin.X, y-origin.Y).RGBA()
				oRatio := float64(oa) / math.MaxUint16

				if oa == 0 {
					ret.(draw.Image).Set(x, y, img.At(x, y))
				} else {
					opacity := utils.LERP(0.0, float64(oa)/math.MaxUint16, effect.opacity)

					if opacity == math.MaxUint16 {
						ret.(draw.Image).Set(x, y, color.RGBA64{R: uint16(or), G: uint16(og), B: uint16(ob)})
					} else {
						r, g, b, a := img.At(x, y).RGBA()
						ratio := float64(a) / math.MaxUint16

						newRed := utils.LERP(float64(or)*oRatio, float64(r)*ratio, opacity)
						newGreen := utils.LERP(float64(og)*oRatio, float64(g)*ratio, opacity)
						newBlue := utils.LERP(float64(ob)*oRatio, float64(b)*ratio, opacity)
						newAlpha := utils.LERP(float64(a),
							math.MaxUint16,
							utils.LERP(0, float64(oa), opacity)/math.MaxUint16)

						ret.(draw.Image).Set(x, y, color.RGBA64{
							R: uint16(utils.ClampUint16(newRed)),
							G: uint16(utils.ClampUint16(newGreen)),
							B: uint16(utils.ClampUint16(newBlue)),
							A: uint16(utils.ClampUint16(newAlpha)),
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

func NewOverlay(stamp image.Image, origin position.Position, opacity float64) *Overlay {
	o := new(Overlay)
	o.stamp = stamp
	o.origin = origin
	o.opacity = opacity

	return o
}
