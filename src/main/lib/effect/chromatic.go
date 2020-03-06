package effect

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"

	core ".."
	utils "../utils"
)

// ChromaticTransformation is an interface that defines the processing
// of information of a single pixel of the image, allowing to create a
// great range of chromatic operations with minimum code duplication
type ChromaticTransformation interface {
	Transform(color.Color) color.Color
}

// GrayscaleTransformation encapsulates the data for a simple grayscale effect
type GrayscaleTransformation struct{}

func (t GrayscaleTransformation) Transform(in color.Color) color.Color {
	return color.Gray16Model.Convert(in)
}

// GrayscaleRatioTransformation encapsulates the data and logic for a
// grayscale effect that uses custom ratios for each color channel
// in order give control on image brightness
type GrayscaleRatioTransformation struct {
	redRatio   float64
	greenRatio float64
	blueRatio  float64
}

func (t GrayscaleRatioTransformation) Transform(in color.Color) color.Color {
	r, g, b, a := in.RGBA()
	gray := utils.ClampUint16(math.Floor(float64(r)*t.redRatio + float64(g)*t.greenRatio + float64(b)*t.blueRatio))
	return color.RGBA64{R: uint16(gray), G: uint16(gray), B: uint16(gray), A: uint16(a)}
}

// SepiaTransformation encapsulates the logic for application of
// a sepia chromatic effect on the color channels of an image
type SepiaTransformation struct{}

func (t SepiaTransformation) Transform(in color.Color) color.Color {
	r, g, b, a := in.RGBA()
	newRed := utils.ClampUint16(math.Floor(float64(r)*.393 + float64(g)*.769 + float64(b)*.189))
	newGreen := utils.ClampUint16(math.Floor(float64(r)*.349 + float64(g)*.686 + float64(b)*.168))
	newBlue := utils.ClampUint16(math.Floor(float64(r)*.272 + float64(g)*.534 + float64(b)*.131))

	return color.RGBA64{R: uint16(newRed), G: uint16(newGreen), B: uint16(newBlue), A: uint16(a)}
}

// NegativeTransformation encapsulates the logic for application of
// a negative chromatic effect on the color channels of an image
type NegativeTransformation struct{}

func (t NegativeTransformation) Transform(in color.Color) color.Color {
	r, g, b, a := in.RGBA()
	newRed := utils.MaxUint16 - r
	newGreen := utils.MaxUint16 - g
	newBlue := utils.MaxUint16 - b

	return color.RGBA64{R: uint16(newRed), G: uint16(newGreen), B: uint16(newBlue), A: uint16(a)}
}

// GenericChromatic serves as a generic customizable structure that encapsulates
// the logic needed to apply a chromatic transformation on a pixel of an image
type GenericChromatic struct {
	core.BaseEffect
	trans ChromaticTransformation
}

func (effect *GenericChromatic) Apply(img image.Image) core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				ret.(draw.Image).Set(x, y, effect.trans.Transform(img.At(x, y)))
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, &contract)
}
