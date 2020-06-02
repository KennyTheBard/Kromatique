package histogram

import (
	"image"
	"image/draw"

	"../utils"
)

func Matching(img image.Image, histogram *Histogram, target []uint, shift ColorShift) image.Image {
	values := histogram.Values()
	cumulative := histogram.Cumulative()
	bounds := img.Bounds()
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
			pxColor := img.At(x, y)
			val := histogram.Eval(pxColor)

			if newVal, ok := mappers[val]; ok {
				pxColor = shift(pxColor, int(newVal)-int(val))
			}

			ret.(draw.Image).Set(x, y, pxColor)
		}
	}

	return ret
}
