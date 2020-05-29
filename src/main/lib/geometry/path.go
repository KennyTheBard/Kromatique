package geometry

import (
	"../utils"
	"math"
)

// Path is an interface that encapsulates any kind of path
type Path interface {
	// GetPoint returns the point found at the given ratio
	// between the two ends of the path
	GetPoint(float64) utils.Point2D
	// GetInterval returns the interval of values for which the
	// points on the curve should be calculated
	GetInterval() (float64, float64)
	// ShouldDivide returns false if 2 consecutive points on the curve
	// calculated for the 2 given values are close enough for the path
	// between them can be approximated through linear interpolation;
	// if true, than there should be at least one or more points calculated
	// between the 2 in order to have the approximated curve be as true
	// to the original one as possible (and computational efficient)
	ShouldDivide(float64, float64) bool
}

// Segment encapsulates a simple straight path defined between 2 points
type Segment struct {
	start, end utils.Point2D
}

func (path *Segment) GetPoint(t float64) utils.Point2D {
	return utils.Pt2D(
		(1-t)*path.start.X+t*path.end.X,
		(1-t)*path.start.Y+t*path.end.Y)
}

func (path *Segment) GetInterval() (float64, float64) {
	return 0, 1
}

func (path *Segment) ShouldDivide(t1, t2 float64) bool {
	return t1 == t2
}

func NewSegment(start, end utils.Point2D) *Segment {
	path := new(Segment)
	path.start = start
	path.end = end

	return path
}

// ComposedPath encapsulates a path composed of multiple independent ones
type ComposedPath struct {
	subpaths []Path
}

func (path *ComposedPath) GetPoint(t float64) utils.Point2D {
	aux := t
	if t == float64(len(path.subpaths)) {
		aux -= 0.0001
	}
	return path.subpaths[int(math.Floor(aux))].GetPoint(t)
}

func (path *ComposedPath) GetInterval() (float64, float64) {
	return 0, float64(len(path.subpaths))
}

func (path *ComposedPath) ShouldDivide(t1, t2 float64) bool {
	// the arguments must be in increasing order
	if t1 > t2 {
		t1, t2 = t2, t1
	}

	// if the points are from different subpaths
	ft1, ft2 := math.Floor(t1), math.Floor(t2)
	if ft1 != ft2 {
		rt1, rt2 := math.Round(t1), math.Round(t2)
		if rt1 != rt2 {
			return true
		}

		return path.subpaths[int(ft1)].ShouldDivide(t1-math.Abs(rt1-t1), t1) ||
			path.subpaths[int(ft2)].ShouldDivide(t2, t2+math.Abs(rt2-t2))
	}

	return path.subpaths[int(ft1)].ShouldDivide(t1, t2)
}

func NewComposedPath(subpaths []Path) *ComposedPath {
	path := new(ComposedPath)
	path.subpaths = subpaths

	return path
}
