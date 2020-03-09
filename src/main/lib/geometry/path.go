package geometry

import (
	"../utils"
)

type Path interface {
	GetPoint(float64) *utils.Point2D
	GetStart() *utils.Point2D
	GetEnd() *utils.Point2D
}

type Segment struct {
	start, end *utils.Point2D
}

func (s Segment) GetPoint(t float64) *utils.Point2D {
	return utils.NewPoint2D(
		(1-t)*s.start.X+t*s.end.X,
		(1-t)*s.start.Y+t*s.end.Y)
}

func (s Segment) GetStart() *utils.Point2D {
	return s.start
}

func (s Segment) GetEnd() *utils.Point2D {
	return s.end
}

func NewSegment(start, end *utils.Point2D) *Segment {
	s := new(Segment)
	s.start = start
	s.end = end

	return s
}
