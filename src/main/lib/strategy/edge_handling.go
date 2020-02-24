package strategy

import (
	"image"
	"image/color"
)

type EdgeHandlingStrategy func(*image.Image, int, int) color.Color

func Extend(img *image.Image, x, y int) color.Color {
	bounds := (*img).Bounds()

	if x < bounds.Min.X {
		x = bounds.Min.X
	} else if x > bounds.Max.X {
		x = bounds.Max.X
	}

	if y < bounds.Min.Y {
		y = bounds.Min.Y
	} else if y > bounds.Max.Y {
		y = bounds.Max.Y
	}

	return (*img).At(x, y)
}

func Wrap(img *image.Image, x, y int) color.Color {
	bounds := (*img).Bounds()

	if x < bounds.Min.X {
		x += bounds.Max.X - bounds.Min.X
	} else if x > bounds.Max.X {
		x -= bounds.Max.X - bounds.Min.X
	}

	if y < bounds.Min.Y {
		y += bounds.Max.Y - bounds.Min.Y
	} else if y > bounds.Max.Y {
		y -= bounds.Max.Y - bounds.Min.Y
	}

	return (*img).At(x, y)
}

func Mirror(img *image.Image, x, y int) color.Color {
	bounds := (*img).Bounds()

	if x < bounds.Min.X {
		x = 2 * bounds.Min.X - x
	} else if x > bounds.Max.X {
		x = 2 * bounds.Max.X - x
	}

	if y < bounds.Min.Y {
		y = 2 * bounds.Min.Y - y
	} else if y > bounds.Max.Y {
		y = 2 * bounds.Max.Y - y
	}

	return (*img).At(x, y)
}

