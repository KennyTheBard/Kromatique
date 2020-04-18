package geometry

import (
	"../utils"
)

// Shape is an interface that encapsulates any kind of 2D shape
type Shape interface {
	// Inside returns if the shape contains the given point
	Inside(utils.Point2D) bool
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
