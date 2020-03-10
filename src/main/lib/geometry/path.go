package geometry

import (
	"../utils"
)

// Path is an interface that encapsulates any kind of path
type Path interface {
	// GetPoint returns the point found at the given ratio
	// between the two ends of the path
	GetPoint(float64) *utils.Point2D
}

// Segment encapsulates a simple straight path defined between 2 points
type Segment struct {
	start, end *utils.Point2D
}

func (s Segment) GetPoint(t float64) *utils.Point2D {
	return utils.NewPoint2D(
		(1-t)*s.start.X+t*s.end.X,
		(1-t)*s.start.Y+t*s.end.Y)
}

func NewSegment(start, end *utils.Point2D) *Segment {
	s := new(Segment)
	s.start = start
	s.end = end

	return s
}
