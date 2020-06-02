package histogram

import (
	"../utils"
	"image"
)

// Histogram encapsulates the data extracted from an image bundled with the logic
// used to extract it in order to apply meaningful transformations to it
type Histogram struct {
	Eval utils.ColorEvaluation

	values    []uint
	numValues uint
}

func (h *Histogram) Scan(img image.Image) {
	h.values = make([]uint, h.numValues)

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			idx := h.Eval(img.At(x, y))
			if idx < h.numValues {
				h.values[idx] += 1
			}
		}
	}
}

func (h *Histogram) Values() []uint {
	return h.values
}

func (h *Histogram) Cumulative() []uint {
	cumulative := make([]uint, len(h.values))
	if len(h.values) > 0 {
		cumulative[0] = h.values[0]
		for idx := 1; idx < len(h.values); idx++ {
			cumulative[idx] = h.values[idx] + cumulative[idx-1]
		}
	}

	return cumulative
}

func NewHistogram(evaluate utils.ColorEvaluation, numValues uint) *Histogram {
	histogram := new(Histogram)
	histogram.Eval = evaluate
	histogram.numValues = numValues

	return histogram
}
