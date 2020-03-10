package render

import (
	"../utils"
	"image/color"
)

// Renderer encapsulates the logic used to draw along a path
type Renderer interface {
	Width() int
	Render(int) color.Color
}

// MatteRenderer encapsulates logic and data to draw a simple matte line
type MatteRenderer struct {
	width int
	color color.Color
}

func (r MatteRenderer) Width() int {
	return r.width
}

func (r MatteRenderer) Render(proximity int) color.Color {
	if proximity != 0 {
		return r.color
	}

	return color.Transparent
}

func NewMatteRenderer(width int, color color.Color) *MatteRenderer {
	r := new(MatteRenderer)
	r.width = width
	r.color = color

	return r
}

// SprayRenderer encapsulates logic and data to draw a simple matte line
type SprayRenderer struct {
	width int
	color color.Color
}

func (sr SprayRenderer) Width() int {
	return sr.width
}

func (sr SprayRenderer) Render(proximity int) color.Color {
	if proximity == 0 {
		return color.Transparent
	}

	r, g, b, _ := sr.color.RGBA()
	return utils.PixelLERP(
		color.RGBA64{
			R: uint16(r),
			G: uint16(g),
			B: uint16(b),
			A: 0},
		sr.color,
		float64(proximity)/float64(sr.width))
}

func NewSprayRenderer(width int, color color.Color) *SprayRenderer {
	sr := new(SprayRenderer)
	sr.width = width
	sr.color = color

	return sr
}
