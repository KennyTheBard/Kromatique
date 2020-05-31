package geometry

import (
	"../utils"
	"image"
	"math"
)

// Shape is an interface that encapsulates any kind of 2D shape
type Shape interface {
	IObject2D
	// Inside returns if the shape contains the given point
	Inside(Point2D) bool
	// MBR = Minimum Bounding Rectangle returns the smallest rectangle
	// that contains the entire shape; should be used to minimize the
	// Inside method calls
	MBR() image.Rectangle
}

// Circle encapsulates a simple circle shape defined by
// a center and a radius
type Circle struct {
	center Point2D
	radius float64
}

func (shape *Circle) Translate(x, y float64) {
	shape.center.Translate(x, y)
}

func (shape *Circle) Scale(x, y float64) {
	shape.center.Scale(x, y)
}

func (shape *Circle) Rotate(a float64) {
	shape.center.Rotate(a)
}

func (shape Circle) Inside(p Point2D) bool {
	return shape.center.Dist(p) <= shape.radius
}

func (shape Circle) MBR() image.Rectangle {
	return image.Rect(
		int(math.Floor(shape.center.X-shape.radius)),
		int(math.Floor(shape.center.Y-shape.radius)),
		int(math.Ceil(shape.center.X+shape.radius)),
		int(math.Ceil(shape.center.Y+shape.radius)))
}

// Rectangle encapsulates a simple rectangle defined by
// an encapsulated image.Rectangle object
type Rectangle struct {
	rect image.Rectangle
}

func (s Rectangle) Inside(p Point2D) bool {
	return p.In(s.rect)
}

func (s Rectangle) MBR() image.Rectangle {
	return s.rect
}
