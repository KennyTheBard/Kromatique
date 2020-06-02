package render

import "image/color"

type Painter func(int, int) color.Color

func MattePainter(c color.Color) Painter {
	return func(_, _ int) color.Color {
		return c
	}
}
