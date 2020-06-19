package histogram

import (
	"../utils"
	"image"
	"image/draw"
	"math"

	"../strategy"
)

type Histogram interface {
	Values() [256]uint
	Cumulative() [256]uint
}

type ValueHistogram struct {
	values [256]uint
}

func (h *ValueHistogram) Values() [256]uint {
	return h.values
}

func (h *ValueHistogram) Cumulative() [256]uint {
	var cumulative [256]uint
	cumulative[0] = h.values[0]
	for idx := 1; idx < len(h.values); idx++ {
		cumulative[idx] = h.values[idx] + cumulative[idx-1]
	}

	return cumulative
}

func NewValueHistogram(values [256]uint) *ValueHistogram {
	h := new(ValueHistogram)
	h.values = values

	return h
}

// ImageHistogram encapsulates the data extracted from an image bundled with the logic
// used to extract it in order to apply meaningful transformations to it
type ImageHistogram struct {
	eval     strategy.ColorEvaluation
	original image.Image
	values   [256]uint
}

func (h *ImageHistogram) Values() [256]uint {
	return h.values
}

func (h *ImageHistogram) Cumulative() [256]uint {
	var cumulative [256]uint
	cumulative[0] = h.values[0]
	for idx := 1; idx < len(h.values); idx++ {
		cumulative[idx] = h.values[idx] + cumulative[idx-1]
	}

	return cumulative
}

// Equalize returns a new image corresponding to the last image scanned
// with this histogram, having a cumulative histogram as close to
// a linear ramp as possible with the available values
func (h *ImageHistogram) Equalize(correction strategy.ColorCorrection) image.Image {
	values := h.Values()
	var idealValues [256]uint

	total := 0
	for _, value := range values {
		total += int(value)
	}
	individualValue := uint(math.Round(float64(total) / (math.MaxUint8 + 1)))

	// create a mapper for each interval
	for idx := range values {
		idealValues[idx] = individualValue
	}

	return h.Match(NewValueHistogram(idealValues), correction)
}

// Match returns a new image corresponding to the last image scanned
// with this histogram, having a cumulative histogram as close to
// a given set of values as possible with the available values
func (h *ImageHistogram) Match(target Histogram, correction strategy.ColorCorrection) image.Image {
	values := h.Values()
	targetCumulative := target.Cumulative()
	//cumulative := h.Cumulative()
	bounds := h.original.Bounds()
	valueMap := make([][]int, bounds.Dy())
	for i := 0; i < bounds.Dy(); i++ {
		valueMap[i] = make([]int, bounds.Dx())
	}

	//// calculate the length of the value space
	//space := uint(len(values))
	//
	//// calculate slope of ideal cumulative
	//slope := float64(cumulative[len(cumulative)-1]) / float64(space)

	// create a mapper for each interval
	mappers := make(map[uint8]uint8)
	end := 0
	var realCumulative uint
	for idx, y := range values {
		if y == 0 {
			continue
		}

		realCumulative += y
		var intervalLen, nextEnd int
		if idx < math.MaxUint8 {
			// in order to overcome slope = 0 problems
			for targetCumulative[end+1] == targetCumulative[end] {
				end += 1
			}
			slope := float64(targetCumulative[end+1] - targetCumulative[end])
			intervalLen = int(utils.ClampUint32((float64(realCumulative) - float64(targetCumulative[end])) / slope))

			nextEnd = end + intervalLen
		}

		mappers[uint8(idx)] = uint8(end)

		end = nextEnd
	}

	// create new image with equalized color
	ret := utils.CreateRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pxColor := h.original.At(x, y)
			val := h.eval(pxColor)

			if newVal, ok := mappers[val]; ok {
				pxColor = correction(pxColor, int16(newVal)-int16(val))
			}

			ret.(draw.Image).Set(x, y, pxColor)
		}
	}

	return ret
}

func NewImageHistogram(img image.Image, eval strategy.ColorEvaluation) *ImageHistogram {
	h := new(ImageHistogram)
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
