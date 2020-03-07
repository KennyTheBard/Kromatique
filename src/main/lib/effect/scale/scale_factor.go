package scale

import "image"

// ScaleFactor represents the value a point must be multiplied by
// in order to obtain its color coordinates on the source image
type ScaleFactor struct {
	X, Y float64
}

// ToDestination returns the color coordinates of the destination image
func (factor ScaleFactor) ToDestination(x, y float64) (float64, float64) {
	return x * factor.X, y * factor.Y
}

// ToSource returns the color coordinates of the source image
func (factor ScaleFactor) ToSource(x, y float64) (float64, float64) {
	return x / factor.X, y / factor.Y
}

// ScaleFactorStrategy encapsulates a strategy used to obtain the
// scale factor of an image and the target size of the said image
type ScaleFactorStrategy interface {

	// Factor returns the scale factor for the given image
	Factor(image.Rectangle) ScaleFactor

	// Size returns the size of the final image
	Size(image.Rectangle) image.Rectangle
}

// FixedSize is an implementation of the ScaleFactorStrategy that
// always returns the same size, but computes different factors
// depending on the given source image dimensions. Useful in
// order to normalize multiple images of different sizes to a
// common size that can be used for comparison.
type FixedSize struct {
	targetSize image.Rectangle
}

func (s FixedSize) Factor(bounds image.Rectangle) ScaleFactor {
	return ScaleFactor{
		X: float64(bounds.Dx()) / float64(s.targetSize.Dx()),
		Y: float64(bounds.Dy()) / float64(s.targetSize.Dy()),
	}
}

func (s FixedSize) Size(bounds image.Rectangle) image.Rectangle {
	return s.targetSize
}

func NewFixedSize(targetSize image.Rectangle) *FixedSize {
	fs := new(FixedSize)
	fs.targetSize = targetSize

	return fs
}

// FixedScaleFactor is an implementation of the ScaleFactorStrategy that
// always returns teh same factor, but computes different target sizes
// depending on the given source image dimensions. Useful in order to
// scale down by the same factor multiple images with the same resize.
type FixedScaleFactor struct {
	factor ScaleFactor
}

func (s FixedScaleFactor) Factor(bounds image.Rectangle) ScaleFactor {
	return s.factor
}

func (s FixedScaleFactor) Size(bounds image.Rectangle) image.Rectangle {
	return image.Rect(int(float64(bounds.Min.X)*s.factor.X),
		int(float64(bounds.Min.Y)*s.factor.Y),
		int(float64(bounds.Max.X)*s.factor.X),
		int(float64(bounds.Max.Y)*s.factor.Y))
}

func NewFixedScaleFactor(factor ScaleFactor) *FixedScaleFactor {
	fsf := new(FixedScaleFactor)
	fsf.factor = factor

	return fsf
}
