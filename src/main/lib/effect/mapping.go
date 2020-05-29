package effect

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"

	"../core"
	"../utils"
)

// MappingRule modifies a color using contained logic
type MappingRule func(color.Color) color.Color

// ColorMapperCondition filters colors based on contained logic
type ColorMapperCondition func(color.Color) bool

// Condition returns a copy of the MappingRule conditioned with
// the given condition; this can be helpful in order to chain conditions
func ConditionRule(rule MappingRule, condition ColorMapperCondition) MappingRule {
	return func(color color.Color) color.Color {
		if condition(color) {
			return rule(color)
		}

		return color
	}
}

// Grayscale maps the given color to a gray shade using the
// standard library default color converter logic
func Grayscale(in color.Color) color.Color {
	return color.Gray16Model.Convert(in)
}

// GrayscaleRatioFactory returns a grayscale mapper that uses
// the given ration to calculate the resulting shade of gray
func GrayscaleRatioFactory(redRatio, greenRatio, blueRatio float64) MappingRule {
	return func(in color.Color) color.Color {
		r, g, b, a := in.RGBA()
		gray := utils.ClampUint16(math.Floor(float64(r)*redRatio + float64(g)*greenRatio + float64(b)*blueRatio))
		return color.RGBA64{R: uint16(gray), G: uint16(gray), B: uint16(gray), A: uint16(a)}
	}
}

// Sepia maps the given color to a sepia shade
func Sepia(in color.Color) color.Color {
	r, g, b, a := in.RGBA()
	newRed := utils.ClampUint16(math.Floor(float64(r)*.393 + float64(g)*.769 + float64(b)*.189))
	newGreen := utils.ClampUint16(math.Floor(float64(r)*.349 + float64(g)*.686 + float64(b)*.168))
	newBlue := utils.ClampUint16(math.Floor(float64(r)*.272 + float64(g)*.534 + float64(b)*.131))

	return color.RGBA64{R: uint16(newRed), G: uint16(newGreen), B: uint16(newBlue), A: uint16(a)}
}

// Negative maps the given color to its negative
func Negative(in color.Color) color.Color {
	r, g, b, a := in.RGBA()
	newRed := math.MaxUint16 - r
	newGreen := math.MaxUint16 - g
	newBlue := math.MaxUint16 - b

	return color.RGBA64{R: uint16(newRed), G: uint16(newGreen), B: uint16(newBlue), A: uint16(a)}
}

// BlackAndWhite maps the given color to black or white depending on a color evaluation
// function and its output to the color input relative to the given threshold
func BlackAndWhite(threshold uint, evaluation utils.ColorEvaluation) MappingRule {
	return func(in color.Color) color.Color {
		val := evaluation(in)
		if val >= threshold {
			return color.White
		}

		return color.Black
	}
}

// ColorMapper serves as a generic customizable structure that encapsulates
// the logic needed to apply a series of MappingRule on an image
type ColorMapper struct {
	engine core.Engine
	rules  []MappingRule
}

func (effect *ColorMapper) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.engine.Contract(img.Bounds().Dy())

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				currentColor := img.At(x, y)
				for _, rule := range effect.rules {
					currentColor = rule(currentColor)
				}

				ret.(draw.Image).Set(x, y, currentColor)
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}
