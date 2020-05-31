package blend

import (
	"image/color"
	"math"

	"../utils"
)

type BlendingStrategy func(color.Color, color.Color) color.Color

func NormalBlending(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, a2 := fg.RGBA()

	if a1+a2 == 0 {
		return color.Transparent
	}

	factor := float64(a2) / float64(a1+a2)

	return color.RGBA64{
		R: uint16(utils.LERP(float64(r1), float64(r2), factor)),
		G: uint16(utils.LERP(float64(g1), float64(g2), factor)),
		B: uint16(utils.LERP(float64(b1), float64(b2), factor)),
		A: uint16(utils.LERP(float64(a1), float64(color.Opaque.A), float64(a2)/float64(color.Opaque.A))),
	}
}

func DifferenceBlending(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, _ := fg.RGBA()

	return NormalBlending(bg, color.RGBA64{
		R: uint16(utils.Abs(int(r1 - r2))),
		G: uint16(utils.Abs(int(g1 - g2))),
		B: uint16(utils.Abs(int(b1 - b2))),
		A: uint16(a1),
	})
}

func SubtractBlending(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, _ := fg.RGBA()

	return NormalBlending(bg, color.RGBA64{
		R: uint16(utils.ClampUint16(float64(r1 - r2))),
		G: uint16(utils.ClampUint16(float64(g1 - g2))),
		B: uint16(utils.ClampUint16(float64(b1 - b2))),
		A: uint16(a1),
	})
}

func DarkenBlending(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, _ := fg.RGBA()

	return NormalBlending(bg, color.RGBA64{
		R: uint16(utils.Min(int(r1), int(r2))),
		G: uint16(utils.Min(int(g1), int(g2))),
		B: uint16(utils.Min(int(b1), int(b2))),
		A: uint16(a1),
	})
}

func LightenBlending(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, _ := fg.RGBA()

	return NormalBlending(bg, color.RGBA64{
		R: uint16(utils.Max(int(r1), int(r2))),
		G: uint16(utils.Max(int(g1), int(g2))),
		B: uint16(utils.Max(int(b1), int(b2))),
		A: uint16(a1),
	})
}

func LinearBurnBlending(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, _ := fg.RGBA()

	var white uint32 = 1 << 16
	return NormalBlending(bg, color.RGBA64{
		R: uint16(utils.ClampUint16(float64(r1 + r2 - white))),
		G: uint16(utils.ClampUint16(float64(g1 + g2 - white))),
		B: uint16(utils.ClampUint16(float64(b1 + b2 - white))),
		A: uint16(a1),
	})
}

func ExclusionBlending(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, _ := fg.RGBA()

	var gray uint32 = (1 << 16) / 2
	return NormalBlending(bg, color.RGBA64{
		R: uint16(utils.ClampUint16(float64(gray - 2*(r1-gray)*(r2-gray)))),
		G: uint16(utils.ClampUint16(float64(gray - 2*(g1-gray)*(g2-gray)))),
		B: uint16(utils.ClampUint16(float64(gray - 2*(b1-gray)*(b2-gray)))),
		A: uint16(a1),
	})
}

func DivideBlending(bg, fg color.Color) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r2, g2, b2, _ := fg.RGBA()

	return NormalBlending(bg, color.RGBA64{
		R: uint16(divideChannel(r1, r2)),
		G: uint16(divideChannel(g1, g2)),
		B: uint16(divideChannel(b1, b2)),
		A: uint16(a1),
	})
}

func divideChannel(c1, c2 uint32) uint32 {
	if c2 == 0 {
		return 1
	}

	return uint32(math.Round(utils.ClampUint16(float64(c1) / float64(c2))))
}
