package mask

import (
	"../geometry"
	"image"
	"image/color"
)

// Mask is an interface that checks if an image process
// should alter a given pixel of an image
type Mask interface {
	Has(int, int) bool
}

// ShapeMask is an implementation of the Mask interface
// that uses a geometric shape as image process mask
type ShapeMask struct {
	geometry.Shape
}

func (mask *ShapeMask) Has(x, y int) bool {
	return mask.Contains(geometry.Pt2D(float64(x), float64(y)))
}

// BitmapMask is an implementation of the Mask interface
// that uses a grayscale image as image process mask
type BitmapMask struct {
	image.Gray
}

func (mask *BitmapMask) Has(x, y int) bool {
	return mask.At(x, y).(color.Gray).Y > 0
}
