package strategy

import (
	"image/color"
	"math"

	"../utils"
)

// ColorMerger defines an interface for all functions used
// to determine the merged result of multiple applied filters for MultiKernel
type ColorMerger func([]color.Color) color.Color

// SobelMerge defines the rule used to merge results for sobel operator;
// despite sobel using only 2 kernels, the function handles more values
func SobelMerge(results []color.Color) color.Color {
	var sumRed, sumGreen, sumBlue float64
	for _, res := range results {
		r, g, b, _ := res.RGBA()
		sumRed += float64(r * r)
		sumGreen += float64(g * g)
		sumBlue += float64(b * b)
	}

	gray := uint16(utils.ClampUint16(math.Round(math.Sqrt((sumRed + sumGreen + sumBlue) / 3))))
	return color.RGBA64{
		R: gray,
		G: gray,
		B: gray,
		A: math.MaxUint16,
	}
}
