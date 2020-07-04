package render

import (
	"image/color"

	"github.com/kennythebard/kromatique/utils"
)

type Painter func(float64, int, int) color.Color

func MattePainter(c color.Color) Painter {
	return func(t float64, _, _ int) color.Color {
		return utils.PixelLERP(color.Transparent, c, t)
	}
}
