package strategy

import "image"

// EdgeHandling defines an interface for all functions used
// to determine the behaviour of filtering around the edge of the image
type EdgeHandling func(image.Rectangle, int, int) (int, int)

// Extends returns the color of the closest pixel of the image
func Extend(bounds image.Rectangle, x, y int) (int, int) {
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

	return x, y
}

// Wrap returns the color of the pixel as if the image is conceptually
// wrapped (or tiled) and values are taken from the opposite edge or corner.
func Wrap(bounds image.Rectangle, x, y int) (int, int) {
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

	return x, y
}

// Mirror returns the color of the pixel as if the image is conceptually
// mirrored at the edges. For example, attempting to read a pixel 3 units
// outside an edge reads one 3 units inside the edge instead.
func Mirror(bounds image.Rectangle, x, y int) (int, int) {
	if x < bounds.Min.X {
		x = 2*bounds.Min.X - x
	} else if x > bounds.Max.X {
		x = 2*bounds.Max.X - x
	}

	if y < bounds.Min.Y {
		y = 2*bounds.Min.Y - y
	} else if y > bounds.Max.Y {
		y = 2*bounds.Max.Y - y
	}

	return x, y
}
