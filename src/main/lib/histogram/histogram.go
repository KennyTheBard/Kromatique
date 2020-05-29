package histogram

import (
	"../utils"
	"image"
)

// Histogram encapsulates the data extracted from an image bundled with the logic
// used to extract it in order to apply meaningful transformations to it
type Histogram struct {
	Eval utils.ColorEvaluation

	values       []uint
	scannedImage image.Image
	start, end   int
}

func (h *Histogram) Scan(img image.Image) {
	h.scannedImage = img
	h.values = make([]uint, h.end-h.start+1)

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			idx := int(h.Eval(img.At(x, y)))
			if idx >= h.start || idx <= h.end {
				h.values[idx-h.start] += 1
			}
		}
	}
}

func (h *Histogram) Original() image.Image {
	return h.scannedImage
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

func (h *Histogram) ValuesInterval() (int, int) {
	return h.start, h.end
}

func NewHistogram(evaluate utils.ColorEvaluation, start, end int) *Histogram {
	histogram := new(Histogram)
	histogram.Eval = evaluate
	histogram.scannedImage = nil
	histogram.start = start
	histogram.end = end

	return histogram
}
