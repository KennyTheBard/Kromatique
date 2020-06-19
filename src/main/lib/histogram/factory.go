package histogram

import (
	"image"

	"../core"
	"../strategy"
)

type Factory struct {
	engine core.Engine
}

func NewFactory(engine core.Engine) *Factory {
	f := new(Factory)
	f.engine = engine

	return f
}

func (f Factory) ImageHistogram(img image.Image, eval strategy.ColorEvaluation) *ImageHistogram {
	h := new(ImageHistogram)
	h.engine = f.engine
	h.eval = eval
	h.original = img
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			idx := int(h.eval(img.At(x, y)))
			h.values[idx] += 1
		}
	}

	return h
}

func (f Factory) ValueHistogram(values [256]uint) *ValueHistogram {
	h := new(ValueHistogram)
	h.values = values

	return h
}
