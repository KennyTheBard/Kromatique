package mapper

import "image/color"

type ColorMapper func(color.Color) color.Color

func ColorMapperFactory(
	condition ColorMapperCondition,
	transformation ColorMapperTransformation) ColorMapper {

	return func(color color.Color) color.Color {
		if condition(color) {
			return transformation(color)
		}

		return color
	}
}
