package lib

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

type Chromatic interface {
	Apply(image.RGBA) image.RGBA
}

type Grayscale struct {
	engine *KromEngine
}

func (gs *Grayscale) Apply(img image.Image) image.Image {
	fmt.Println(img.Bounds())
	ret := Create(img.Bounds())
	contract := gs.engine.Contract(img.Bounds().Dy())

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
	contract.Deadline()

	return ret
}