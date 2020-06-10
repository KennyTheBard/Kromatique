package histogram

import (
	"../utils"
	"image"
	"image/draw"
	"math"

	"../strategy"
)

// Histogram encapsulates the data extracted from an image bundled with the logic
// used to extract it in order to apply meaningful transformations to it
type Histogram struct {
	eval      strategy.Evaluation
	original  image.Image
	values    []uint
	numValues uint
}

func (h *Histogram) Scan(img image.Image) {
	h.values = make([]uint, h.numValues)

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			idx := h.eval(img.At(x, y))
			if idx < h.numValues {
				h.values[idx] += 1
			}
		}
	}

	h.original = nil
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

// Equalize returns a new image corresponding to the last image scanned
// with this histogram, having a cumulative histogram as close to
// a linear ramp as possible with the available values
func (h *Histogram) Equalize(shift strategy.ColorShift) image.Image {
	values := h.Values()
	cumulative := h.Cumulative()
	idealCumulative := make([]uint, len(values))

	// calculate the length of the value space
	space := uint(len(values))

	// calculate slope of ideal cumulative
	slope := float64(cumulative[len(cumulative)-1]) / float64(space)

	// create a mapper for each interval
	for idx := range values {
		idealCumulative[idx] = uint(math.Round(slope * float64(idx)))
	}

	return h.Match(idealCumulative, shift)
}

// Match returns a new image corresponding to the last image scanned
// with this histogram, having a cumulative histogram as close to
// a given set of values as possible with the available values
func (h *Histogram) Match(target []uint, shift strategy.ColorShift) image.Image {
	values := h.Values()
	cumulative := h.Cumulative()
	bounds := h.original.Bounds()
	valueMap := make([][]int, bounds.Dy())
	for i := 0; i < bounds.Dy(); i++ {
		valueMap[i] = make([]int, bounds.Dx())
	}

	// calculate the length of the value space
	space := uint(len(values))

	// calculate slope of ideal cumulative
	slope := float64(cumulative[len(cumulative)-1]) / float64(space)

	// create a mapper for each interval
	mappers := make(map[uint]uint)
	prevEnd := 0
	var realCumulative uint
	for idx, y := range values {
		if y == 0 {
			continue
		}

		realCumulative += y
		var intervalLen, nextEnd int
		if idx < len(values)-1 {
			intervalLen = int((float64(realCumulative) - float64(target[prevEnd])) / (slope))
			nextEnd = prevEnd + intervalLen
		}

		mappers[uint(idx)] = uint(prevEnd)

		prevEnd = nextEnd
	}

	// create new image with equalized color
	ret := utils.CreateRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pxColor := h.original.At(x, y)
			val := h.eval(pxColor)

			if newVal, ok := mappers[val]; ok {
				pxColor = shift(pxColor, int(newVal)-int(val))
			}

			ret.(draw.Image).Set(x, y, pxColor)
		}
	}

	return ret
}

func NewHistogram(evaluate strategy.Evaluation, numValues uint) *Histogram {
	histogram := new(Histogram)
	histogram.eval = evaluate
	histogram.numValues = numValues
	histogram.original = nil

	return histogram
}
