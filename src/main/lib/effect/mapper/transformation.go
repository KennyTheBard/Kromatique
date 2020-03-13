package mapper

import "image/color"

type ColorMapperTransformation func(color.Color) color.Color
