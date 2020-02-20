package lib

import (
	"image"
	"image/color"
)

type SingleKernelFilter interface {
	FilterSize() int
	Filter
}

type SKFilter struct {
	engine   KromEngine
	strategy EdgeHandlingStrategy
	Matrix   [][]float64
}

func (s *simpleFilter) FilterSize() int {
	return len(s.Matrix)
}

//func (s *simpleFilter) Apply(image.RGBA) image.RGBA {
//
//}
