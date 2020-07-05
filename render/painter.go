package render

import (
	"image/color"

	"github.com/kennythebard/kromatique/utils"
)

// Painter is an interface for a function that
// recieves a set of coordinates and a factor value
// between 0 and 1 and returns a color
type Painter func(float64, int, int) color.Color

// ColorPainter returns a Painter function that returns
// the linear interpolation between color.Transparent
// and the given color, based on the given factor
func ColorPainter(c color.Color) Painter {
	return func(t float64, _, _ int) color.Color {
		return utils.PixelLERP(color.Transparent, c, t)
	}
}
