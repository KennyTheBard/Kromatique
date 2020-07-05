package utils

import (
	"image/color"
	"math"
)

// LERP is a linear interpolation function
func LERP(p1, p2, alpha float64) float64 {
	return (1-alpha)*p1 + alpha*p2
}

// BiLERP is a bilinear interpolation function
func BiLERP(p1, p2, p3, p4, alpha, beta float64) float64 {
	return LERP(LERP(p1, p2, alpha), LERP(p3, p4, alpha), beta)
}

// PixelLERP is an application of LERP on each color channel
func PixelLERP(p1, p2 color.Color, alpha float64) color.Color {
	r1, g1, b1, a1 := p1.RGBA()
	r2, g2, b2, a2 := p2.RGBA()

	return color.RGBA64{
		R: uint16(math.Round(LERP(float64(r1), float64(r2), alpha))),
		G: uint16(math.Round(LERP(float64(g1), float64(g2), alpha))),
		B: uint16(math.Round(LERP(float64(b1), float64(b2), alpha))),
		A: uint16(math.Round(LERP(float64(a1), float64(a2), alpha)))}
}

// PixelBiLERP is an application of BiLERP on each color channel
func PixelBiLERP(p1, p2, p3, p4 color.Color, alpha, beta float64) color.Color {
	return PixelLERP(PixelLERP(p1, p2, alpha), PixelLERP(p3, p4, alpha), beta)
}
