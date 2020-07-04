package mask

import (
	"image"
	"math"
)

// Mask is a function that returns how much a color
// should change for a given coordinate of the image
type Mask func(int, int) float64

func TrueMask(int, int) float64 {
	return 1
}

//// ShapeMask returns a mask based on the given shape
//func ShapeMask(shape geometry.Shape) Mask {
//	return func(x, y int) float64 {
//		if shape.Contains(geometry.Pt2D(float64(x), float64(y))) {
//			return 1
//		} else {
//			return 0
//		}
//	}
//}

// BitmapMask returns a mask based on a bitmap of values
func BitmapMask(bitmap [][]float64) Mask {
	var dx, dy int
	dy = len(bitmap)
	if dy > 0 {
		dx = len(bitmap[0])
	}

	return func(x, y int) float64 {
		if y >= dy || x >= dx {
			return 0
		} else {
			return bitmap[y][x]
		}
	}
}

// GrayMask returns a mask based on a gray image
func GrayMask(img image.Gray) Mask {
	return func(x, y int) float64 {
		return float64(img.GrayAt(x, y).Y) / (math.MaxUint8 - 1)
	}
}
