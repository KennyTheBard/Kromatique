package blend

import (
	"image"
	"image/draw"

	"github.com/kennythebard/kromatique/core"
	"github.com/kennythebard/kromatique/mask"
	"github.com/kennythebard/kromatique/utils"
)

// Blender is a function that receives 2 images, a background and a foreground one
// a position for the origin of the second and a mask to interpolate with
type Blender func(image.Image, image.Image, image.Point, mask.Mask) image.Image

// Blend returns a Blender function using the given BlendingStrategy
func Blend(blendingStrategy BlendingStrategy) Blender {
	return func(bg, fg image.Image, fgOrigin image.Point, fgMask mask.Mask) image.Image {
		ret := utils.CreateRGBA(bg.Bounds())

		core.Parallelize(bg.Bounds().Dy(), func(y int) {
			for x := bg.Bounds().Min.X; x < bg.Bounds().Max.X; x++ {
				cbg, cfg := bg.At(x, y), fg.At(x+fgOrigin.X, y+fgOrigin.Y)

				pxColor := cbg
				if image.Pt(x, y).In(fg.Bounds()) {
					pxColor = utils.PixelLERP(cbg, blendingStrategy(cbg, cfg), fgMask(x+fgOrigin.X, y+fgOrigin.Y))
				}

				ret.(draw.Image).Set(x, y, pxColor)
			}
		})

		return ret
	}
}
