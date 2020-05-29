package histogram

import (
	"../utils"
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// ColorShift is a function that shifts the value of a color by
// a certain degree, be it negative or positive
type ColorShift func(color.Color, int) color.Color

// UniformColorShift modifies all 3 color channels with the entire value
func UniformColorShift(c color.Color, val int) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(utils.ClampUint8(float64(int(r>>8) + val))),
		G: uint8(utils.ClampUint8(float64(int(g>>8) + val))),
		B: uint8(utils.ClampUint8(float64(int(b>>8) + val))),
		A: uint8(a >> 8),
	}
}

// DividedColorShift modifies all 3 color channels by equal part
// of shifting value, clamping the result
func DividedColorShift(c color.Color, val int) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(utils.ClampUint8(float64(int(r>>8) + val/3))),
		G: uint8(utils.ClampUint8(float64(int(g>>8) + val/3))),
		B: uint8(utils.ClampUint8(float64(int(b>>8) + val/3))),
		A: uint8(a >> 8),
	}
}

func Equalize(histogram *Histogram, shift ColorShift) image.Image {
	values := histogram.Values()
	cumulative := histogram.Cumulative()
	start, end := histogram.ValuesInterval()
	img := histogram.Original()
	bounds := img.Bounds()
	valueMap := make([][]int, bounds.Dy())
	for i := 0; i < bounds.Dy(); i++ {
		valueMap[i] = make([]int, bounds.Dx())
	}

	// calculate the length of the value space
	space := uint(end - start)

	// calculate slope of ideal cumulative
	slope := float64(cumulative[len(cumulative)-1]) / float64(space)

	// create a mapper for each interval
	mappers := make(map[uint]uint)
	prevEnd := start
	curr := 0
	var realCumulative uint
	for idx, y := range values {
		if y == 0 {
			continue
		}

		idealCumulative := slope * float64(prevEnd)
		realCumulative += y
		var intervalLen, nextEnd int
		if idx < len(values)-1 {
			intervalLen = int((float64(realCumulative) - idealCumulative) / (slope))
			if intervalLen < 1 {
				intervalLen = 1
			}
			nextEnd = prevEnd + intervalLen
		} else {
			// last interval
			nextEnd = end
		}

		mappers[uint(start+idx)] = uint(start + prevEnd)

		prevEnd = nextEnd
		curr += 1
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

	fmt.Println("mappers", mappers)

	return ret
}
