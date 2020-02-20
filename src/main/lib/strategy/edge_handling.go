package strategy

import (
	"image"
	"image/color"
)

type EdgeHandlingStrategy func(image.Point, *image.Image) color.Color

func Extend(p image.Point, img *image.Image) color.Color {
	bounds := (*img).Bounds()

	if p.X < bounds.Min.X {
		p.X = bounds.Min.X
	} else if p.X > bounds.Max.X {
		p.X = bounds.Max.X
	}

	if p.Y < bounds.Min.Y {
		p.Y = bounds.Min.Y
	} else if p.Y > bounds.Max.Y {
		p.Y = bounds.Max.Y
	}

	return (*img).At(p.X, p.Y)
}

func Wrap(p image.Point, img *image.Image) color.Color {
	bounds := (*img).Bounds()

	if p.X < bounds.Min.X {
		p.X += bounds.Max.X - bounds.Min.X
	} else if p.X > bounds.Max.X {
		p.X -= bounds.Max.X - bounds.Min.X
	}

	if p.Y < bounds.Min.Y {
		p.Y += bounds.Max.Y - bounds.Min.Y
	} else if p.Y > bounds.Max.Y {
		p.Y -= bounds.Max.Y - bounds.Min.Y
	}

	return (*img).At(p.X, p.Y)
}

func Mirror(p image.Point, img *image.Image) color.Color {
	bounds := (*img).Bounds()

	if p.X < bounds.Min.X {
		p.X = 2 * bounds.Min.X - p.X
	} else if p.X > bounds.Max.X {
		p.X = 2 * bounds.Max.X - p.X
	}

	if p.Y < bounds.Min.Y {
		p.Y = 2 * bounds.Min.Y - p.Y
	} else if p.Y > bounds.Max.Y {
		p.Y = 2 * bounds.Max.Y - p.Y
	}

	return (*img).At(p.X, p.Y)
}

