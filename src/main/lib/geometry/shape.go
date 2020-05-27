package geometry

import (
	"../utils"
	"image"
	"math"
)

// Shape is an interface that encapsulates any kind of 2D shape
type Shape interface {
	// Inside returns if the shape contains the given point
	Inside(utils.Point2D) bool
	// MBR = Minimum Bounding Rectangle returns the smallest rectangle
	// that contains the entire shape; should be used to minimize the
	// Inside method calls
	MBR() image.Rectangle
}

// Circle encapsulates a simple circle shape defined by
// a center and a radius
type Circle struct {
	center utils.Point2D
	radius float64
}

func (s Circle) Inside(p utils.Point2D) bool {
	return s.center.Dist(p) <= s.radius
}

func (s Circle) MBR() image.Rectangle {
	return image.Rect(
		int(math.Floor(s.center.X-s.radius)),
		int(math.Floor(s.center.Y-s.radius)),
		int(math.Ceil(s.center.X+s.radius)),
		int(math.Ceil(s.center.Y+s.radius)))
}

// Rectangle encapsulates a simple rectangle defined by
// an encapsulated image.Rectangle object
type Rectangle struct {
	rect image.Rectangle
}

func (s Rectangle) Inside(p utils.Point2D) bool {
	return p.In(s.rect)
}

func (s Rectangle) MBR() image.Rectangle {
	return s.rect
}
