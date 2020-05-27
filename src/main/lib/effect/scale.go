package effect

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"

	core ".."
	"../utils"
)

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

// ColorSamplingStrategy encapsulates a strategy used to decide what color
// the pixel at the given position from the destination image will have
type ColorSamplingStrategy func(image.Image, int, int, ScaleFactor) color.Color

// SinglePixelSampling is an implementation of ColorSamplingStrategy that uses
// the color of a single pixel (the closest to the target) from the source image
func SinglePixelSampling(img image.Image, destX, destY int, factor ScaleFactor) color.Color {
	exactX, exactY := factor.ToSource(float64(destX), float64(destY))

	return img.At(int(math.Round(exactX)), int(math.Round(exactY)))
}

// CornerPixelsSampling is an implementation of ColorSamplingStrategy that uses
// a bilinear interpolation between the corners of the rectangle obtained by scaling
// back a pixel from the destination image to the source image; the scaled rectangle
// can actually be smaller than a single pixel would for factors greater than 1
func CornerPixelsSampling(img image.Image, x, y int, factor ScaleFactor) color.Color {
	exactX, exactY := factor.ToSource(float64(x), float64(y))

	return utils.PixelBiLERP(
		img.At(int(math.Floor(exactX)), int(math.Floor(exactY))),
		img.At(int(math.Ceil(exactX)), int(math.Floor(exactY))),
		img.At(int(math.Floor(exactX)), int(math.Ceil(exactY))),
		img.At(int(math.Ceil(exactX)), int(math.Ceil(exactY))),
		exactX-math.Floor(exactX),
		exactY-math.Floor(exactY))
}

// Scale serves as a generic customizable structure that encapsulates
// the logic needed to apply a scaling transformation on an image
type Scale struct {
	core.Base
	scaleFactorStrategy   ScaleFactorStrategy
	colorSamplingStrategy ColorSamplingStrategy
}

func (effect *Scale) Apply(img image.Image) *core.Promise {
	factor := effect.scaleFactorStrategy.Factor(img.Bounds())
	trgBounds := effect.scaleFactorStrategy.Size(img.Bounds())

	ret := utils.CreateRGBA(trgBounds)
	contract := effect.GetEngine().Contract(ret.Bounds().Dy())

	for i := ret.Bounds().Min.Y; i < ret.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := ret.Bounds().Min.X; x < ret.Bounds().Max.X; x++ {
				col := effect.colorSamplingStrategy(img, x, y, factor)

				ret.(draw.Image).Set(x, y, col)
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}

func NewScale(scaleFactorStrategy ScaleFactorStrategy, colorSamplingStrategy ColorSamplingStrategy) *Scale {
	scale := new(Scale)
	scale.scaleFactorStrategy = scaleFactorStrategy
	scale.colorSamplingStrategy = colorSamplingStrategy

	return scale
}
