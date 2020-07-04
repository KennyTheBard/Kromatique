package strategy

import "image"

// ScaleFactor represents the value a point must be multiplied by
// in order to obtain its color coordinates on the source image
type ScaleFactor struct {
	X, Y float64
}

// GetSourcePx returns the color coordinates of the source image
// for the given color coordinates of the destination image
func (factor ScaleFactor) GetSourcePx(x, y float64) (float64, float64) {
	return x / factor.X, y / factor.Y
}

// ScalingTarget encapsulates a strategy used to obtain the
// scale factor of an image and the target size of the said image
type ScalingTarget func(image.Rectangle) (ScaleFactor, image.Rectangle)
