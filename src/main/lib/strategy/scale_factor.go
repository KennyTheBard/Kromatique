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

// FixedSize returns an implementation of the ScalingTarget that
// always returns the same size, but computes different factors
// depending on the given source image dimensions. Useful in
// order to normalize multiple images of different sizes to a
// common size that can be used for comparison.
func FixedSize(targetSize image.Rectangle) ScalingTarget {
	return func(bounds image.Rectangle) (ScaleFactor, image.Rectangle) {
		return ScaleFactor{
			X: float64(bounds.Dx()) / float64(targetSize.Dx()),
			Y: float64(bounds.Dy()) / float64(targetSize.Dy()),
		}, targetSize
	}
}

// FixedScaleFactor returns an implementation of the ScalingTarget that
// always returns teh same factor, but computes different target sizes
// depending on the given source image dimensions. Useful in order to
// scale down by the same factor multiple images with the same resize.
func FixedScaleFactor(factor ScaleFactor) ScalingTarget {
	return func(bounds image.Rectangle) (ScaleFactor, image.Rectangle) {
		return factor, image.Rect(int(float64(bounds.Min.X)*factor.X),
			int(float64(bounds.Min.Y)*factor.Y),
			int(float64(bounds.Max.X)*factor.X),
			int(float64(bounds.Max.Y)*factor.Y))
	}
}
