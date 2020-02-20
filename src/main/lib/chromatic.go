package lib

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Grayscale encapsulates the data for a grayscale effect that uses custom
// ratios for each color channel in order give control on image brightness
type Grayscale struct {
	engine     *KromEngine
	redRatio   int
	greenRatio int
	blueRatio  int
}

func (gs *Grayscale) Apply(img image.Image) Promise {
	ret := CreateRGBA(img.Bounds())
	contract := gs.engine.Contract(img.Bounds().Dy())

	totalRatio := float64(gs.redRatio + gs.greenRatio + gs.blueRatio)
	rr := float64(gs.redRatio) / totalRatio
	gr := float64(gs.greenRatio) / totalRatio
	br := float64(gs.blueRatio) / totalRatio

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if totalRatio == 0 {
			if err := contract.PlaceOrder(func() {
				for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
					ret.(draw.Image).Set(x, y, color.Gray16Model.Convert(img.At(x, y)))
				}
			}); err != nil {
				fmt.Print(err)
				break
			}
		} else {
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
	}

	return Promise{img: ret, contract: &contract}
}

// Sepia encapsulates the data for a sepia effect
type Sepia struct {
	engine *KromEngine
}

func (gs *Sepia) Apply(img image.Image) Promise {
	ret := CreateRGBA(img.Bounds())
	contract := gs.engine.Contract(img.Bounds().Dy())

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				newRed := ClampUint16(math.Floor(float64(r)*.393 + float64(g)*.769 + float64(b)*.189))
				newGreen := ClampUint16(math.Floor(float64(r)*.349 + float64(g)*.686 + float64(b)*.168))
				newBlue := ClampUint16(math.Floor(float64(r)*.272 + float64(g)*.534 + float64(b)*.131))

				ret.(draw.Image).Set(x, y, color.RGBA64{R: uint16(newRed), G: uint16(newGreen), B: uint16(newBlue)})
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return Promise{img: ret, contract: &contract}
}
