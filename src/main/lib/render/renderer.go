package render

import "image/color"

type Renderer interface {
	Width() int
	Render(int) color.Color
}

type SimpleLineRenderer struct {
	width int
	color color.Color
}

func (r SimpleLineRenderer) Width() int {
	return r.width
}

func (r SimpleLineRenderer) Render(proximity int) color.Color {
	if proximity != 0 {
		return r.color
	}

	return color.Transparent
}

func NewSimpleLineRenderer(width int, color color.Color) *SimpleLineRenderer {
	r := new(SimpleLineRenderer)
	r.width = width
	r.color = color

	return r
}
