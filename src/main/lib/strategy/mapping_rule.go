package strategy

import (
	"image/color"
	"math"

	"../utils"
)

// MappingRule modifies a color using contained logic
type MappingRule func(color.Color) color.Color

// Condition filters colors based on contained logic
type Condition func(color.Color) bool

// AddCondition returns MappingRule that execute the given
// MappingRule only if the given Condition evaluates to true
func AddCondition(rule MappingRule, condition Condition) MappingRule {
	return func(color color.Color) color.Color {
		if condition(color) {
			return rule(color)
		}

		return color
	}
}

// ComposeRule returns a MappingRule that executes the first MappingRule if
// the given Condition returns true, or second MappingRule for false
func ComposeRule(condition Condition, ruleTrue, ruleFalse MappingRule) MappingRule {
	return func(color color.Color) color.Color {
		if condition(color) {
			return ruleTrue(color)
		} else {
			return ruleFalse(color)
		}
	}
}

// Identity returns the given color
func Identity(in color.Color) color.Color {
	return in
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
func BlackAndWhite(evaluation ColorEvaluation, threshold uint8) MappingRule {
	return func(in color.Color) color.Color {
		val := evaluation(in)
		if val >= threshold {
			return color.White
		}

		return color.Black
	}
}

// CorrectionMapping maps the given color to a color obtained by applying
// the correction strategy with a given factor on each color pixel
func CorrectionMapping(correction ColorCorrection, factor int16) MappingRule {
	return func(in color.Color) color.Color {
		return correction(in, factor)
	}
}

// BrightnessMapping maps the given color to a brighter or darker color
// depending on the given brightness factor
func BrightnessMapping(brightnessFactor float64) MappingRule {
	return CorrectionMapping(LightnessCorrection, int16(utils.ClampUint8(math.Round(brightnessFactor*(math.MaxInt16+1)))))
}
