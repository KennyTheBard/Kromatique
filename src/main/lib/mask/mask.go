package mask

import (
	"../geometry"
	"image"
	"image/color"
)

type Mask interface {
	Has(int, int) bool
}

type ShapeMask struct {
	shape geometry.Shape
}

func (mask *ShapeMask) Has(x, y int) bool {
	return mask.shape.Inside(geometry.Pt2D(float64(x), float64(y)))
}

type BitmapMask struct {
	bitmap image.Gray
}

func (mask *BitmapMask) Has(x, y int) bool {
	return mask.bitmap.At(x, y).(color.Gray).Y > 0
}
