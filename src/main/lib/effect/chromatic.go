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

// Grayscale encapsulates the data for a simple grayscale effect
type Grayscale struct {
	core.BaseEffect
}

func (effect *Grayscale) Apply(img image.Image) core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				ret.(draw.Image).Set(x, y, color.Gray16Model.Convert(img.At(x, y)))
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, &contract)
}

// RatioGrayscale encapsulates the data for a grayscale effect that uses custom
// ratios for each color channel in order give control on image brightness
type RatioGrayscale struct {
	core.BaseEffect
	redRatio   int
	greenRatio int
	blueRatio  int
}

func (effect *RatioGrayscale) Apply(img image.Image) core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())

	totalRatio := float64(effect.redRatio + effect.greenRatio + effect.blueRatio)
	rr := float64(effect.redRatio) / totalRatio
	gr := float64(effect.greenRatio) / totalRatio
	br := float64(effect.blueRatio) / totalRatio

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				gray := math.Floor(float64(r)*rr + float64(g)*gr + float64(b)*br)
				ret.(draw.Image).Set(x, y, color.Gray16{Y: uint16(gray)})
			}
		}); err != nil {
			fmt.Print(err)
			break
		}

	}

	return core.NewPromise(ret, &contract)
}

// Sepia encapsulates the data for a sepia effect
type Sepia struct {
	core.BaseEffect
}

func (effect *Sepia) Apply(img image.Image) core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				newRed := utils.ClampUint16(math.Floor(float64(r)*.393 + float64(g)*.769 + float64(b)*.189))
				newGreen := utils.ClampUint16(math.Floor(float64(r)*.349 + float64(g)*.686 + float64(b)*.168))
				newBlue := utils.ClampUint16(math.Floor(float64(r)*.272 + float64(g)*.534 + float64(b)*.131))

				ret.(draw.Image).Set(x, y, color.RGBA64{R: uint16(newRed), G: uint16(newGreen), B: uint16(newBlue)})
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, &contract)
}

// Negative encapsulates the data for a negative effect
type Negative struct {
	core.BaseEffect
}

func (effect *Negative) Apply(img image.Image) core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.GetEngine().Contract(img.Bounds().Dy())

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				newRed := utils.MaxUint16 - r
				newGreen := utils.MaxUint16 - g
				newBlue := utils.MaxUint16 - b

				ret.(draw.Image).Set(x, y, color.RGBA64{R: uint16(newRed), G: uint16(newGreen), B: uint16(newBlue)})
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, &contract)
}
