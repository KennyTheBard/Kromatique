package blend

import (
	"image/color"

	"github.com/kennythebard/kromatique/utils"
)

// BlendingStrategy receives two colors and returns a blend between the two
type BlendingStrategy func(color.Color, color.Color) color.Color

// Normal returns a blending between the two colors based on the alpha channel
// of the foreground color
func Normal(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, a2 := fg.RGBA()

	if a1+a2 == 0 {
		return color.Transparent
	}

	return color.RGBA64{
		R: uint16(utils.LERP(float64(r1), float64(r2), float64(a2)/float64(color.Opaque.A))),
		G: uint16(utils.LERP(float64(g1), float64(g2), float64(a2)/float64(color.Opaque.A))),
		B: uint16(utils.LERP(float64(b1), float64(b2), float64(a2)/float64(color.Opaque.A))),
		A: uint16(utils.LERP(float64(a1), float64(color.Opaque.A), float64(a2)/float64(color.Opaque.A))),
	}
}

// Subtract return a blending between the two colors that makes a difference
// between the two before applying a normal blend
func Subtract(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, _ := fg.RGBA()

	return Normal(bg, color.RGBA64{
		R: uint16(utils.ClampUint16(float64(r1 - r2))),
		G: uint16(utils.ClampUint16(float64(g1 - g2))),
		B: uint16(utils.ClampUint16(float64(b1 - b2))),
		A: uint16(a1),
	})
}

// Subtract return a blending between the two colors that uses the lower
// values of each channel of the two colors before applying a normal blend
func Darken(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, _ := fg.RGBA()

	return Normal(bg, color.RGBA64{
		R: uint16(utils.Min(int(r1), int(r2))),
		G: uint16(utils.Min(int(g1), int(g2))),
		B: uint16(utils.Min(int(b1), int(b2))),
		A: uint16(a1),
	})
}

// Subtract return a blending between the two colors that uses the higher
// values of each channel of the two colors before applying a normal blend
func Lighten(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, _ := fg.RGBA()

	return Normal(bg, color.RGBA64{
		R: uint16(utils.Max(int(r1), int(r2))),
		G: uint16(utils.Max(int(g1), int(g2))),
		B: uint16(utils.Max(int(b1), int(b2))),
		A: uint16(a1),
	})
}
