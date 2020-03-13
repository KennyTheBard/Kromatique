package mapper

import "image/color"

type ColorMapperCondition func(color.Color) bool
