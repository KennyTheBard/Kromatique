package lib

import "image"

// Filter is a simple interface over the concept of image filter
// encapsulating it's internal parameters and providing
// a straight forward way to apply it to an image
type Filter interface {
	Apply(image.RGBA) image.RGBA
}


const (
	Extend = iota
	Wrap = iota
	Mirror = iota
	Crop = iota
	KernelCrop = iota
)

type filter struct {
	EdgeHandlingStrategy int
	engine KromEngine
}

type SimpleFilter interface {
	FilterSize() int
	Filter
}

type simpleFilter struct {
	Matrix [][]float64
	filter
}

func (s *simpleFilter) FilterSize() int {
	return len(s.Matrix)
}

//func (s *simpleFilter) Apply(image.RGBA) image.RGBA {
//
//}
