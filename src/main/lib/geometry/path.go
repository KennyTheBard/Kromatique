package geometry

import (
	"math"
)

// Path is an interface that encapsulates any kind of path
type Path interface {
	IObject2D
	// GetPoint returns the point found at the given ratio
	// between the two ends of the path
	GetPoint(float64) Point2D
}

// Segment encapsulates a simple straight path defined between 2 points
type Segment struct {
	Object2D
	p0, p1 Point2D
}

func (path *Segment) GetPoint(t float64) Point2D {
	p0, p1 := path.Model().Apply(path.p0), path.Model().Apply(path.p1)

	return Pt2D(
		(1-t)*p0.X+t*p1.X,
		(1-t)*p0.Y+t*p1.Y)
}

func NewSegment(p0, p1 Point2D) *Segment {
	path := new(Segment)
	InitObject(&path.Object2D)
	path.p0 = p0
	path.p1 = p1

	return path
}

// Bezier2 implements a linear Bezier curve
type Bezier2 struct {
	Object2D
	p0, p1 Point2D
}

func (path *Bezier2) GetPoint(t float64) Point2D {
	p0, p1 := path.Model().Apply(path.p0), path.Model().Apply(path.p1)

	return Pt2D(
		(1-t)*p0.X+t*p1.X,
		(1-t)*p0.Y+t*p1.Y)
}

func NewBezier2(p0, p1 Point2D) *Bezier2 {
	path := new(Bezier2)
	InitObject(&path.Object2D)
	path.p0 = p0
	path.p1 = p1

	return path
}

// Bezier3 implements a quadratic Bezier curve
type Bezier3 struct {
	Object2D
	p0, p1, p2 Point2D
}

func (path *Bezier3) GetPoint(t float64) Point2D {
	p0, p1, p2 := path.Model().Apply(path.p0), path.Model().Apply(path.p1), path.Model().Apply(path.p2)

	return Pt2D(
		(1-t)*(1-t)*p0.X+2*t*(1-t)*p1.X+t*t*p2.X,
		(1-t)*(1-t)*p0.Y+2*t*(1-t)*p1.Y+t*t*p2.Y)
}

func NewBezier3(p0, p1, p2 Point2D) *Bezier3 {
	path := new(Bezier3)
	InitObject(&path.Object2D)
	path.p0 = p0
	path.p1 = p1
	path.p2 = p2

	return path
}

// Bezier4 implements a cubic Bezier curve
type Bezier4 struct {
	Object2D
	p0, p1, p2, p3 Point2D
}

func (path *Bezier4) GetPoint(t float64) Point2D {
	p0, p1, p2, p3 := path.Model().Apply(path.p0), path.Model().Apply(path.p1), path.Model().Apply(path.p2), path.Model().Apply(path.p3)

	return Pt2D(
		(1-t)*(1-t)*(1-t)*p0.X+3*t*(1-t)*(1-t)*p1.X+3*t*t*(1-t)*p2.X+t*t*t*p3.X,
		(1-t)*(1-t)*(1-t)*p0.Y+3*t*(1-t)*(1-t)*p1.Y+3*t*t*(1-t)*p2.Y+t*t*t*p3.Y)
}

func NewBezier4(p0, p1, p2, p3 Point2D) *Bezier4 {
	path := new(Bezier4)
	InitObject(&path.Object2D)
	path.p0 = p0
	path.p1 = p1
	path.p2 = p2
	path.p3 = p3

	return path
}

// Hermite implements a cubic Hermite curve
type Hermite struct {
	Object2D
	p0, p1, m0, m1 Point2D
}

func (path *Hermite) GetPoint(t float64) Point2D {
	p0, p1, m0, m1 := path.Model().Apply(path.p0), path.Model().Apply(path.p1), path.Model().Apply(path.m0), path.Model().Apply(path.m1)

	return Pt2D(
		(2*t*t*t-3*t*t+1)*p0.X+(t*t*t-2*t*t+t)*m0.X+(-2*t*t*t+3*t*t)*p1.X+(t*t*t-t*t)*m1.X,
		(2*t*t*t-3*t*t+1)*p0.Y+(t*t*t-2*t*t+t)*m0.Y+(-2*t*t*t+3*t*t)*p1.Y+(t*t*t-t*t)*m1.Y)
}

func NewHermite(p0, p1, m0, m1 Point2D) *Hermite {
	path := new(Hermite)
	InitObject(&path.Object2D)
	path.p0 = p0
	path.p1 = p1
	path.m0 = m0
	path.m1 = m1

	return path
}

// ComposedPath encapsulates a path composed of multiple independent ones
type ComposedPath struct {
	subpaths []Path
}

func (path *ComposedPath) Translate(x, y float64) {
	for _, subpath := range path.subpaths {
		subpath.Translate(x, y)
	}
}

func (path *ComposedPath) Scale(x, y float64) {
	for _, subpath := range path.subpaths {
		subpath.Scale(x, y)
	}
}

func (path *ComposedPath) Rotate(a float64) {
	for _, subpath := range path.subpaths {
		subpath.Rotate(a)
	}
}

func (path *ComposedPath) GetPoint(t float64) Point2D {
	aux := t
	if t == float64(len(path.subpaths)) {
		aux -= 0.0001
	}
	return path.subpaths[int(math.Floor(aux))].GetPoint(t)
}

func NewComposedPath(subpaths []Path) *ComposedPath {
	path := new(ComposedPath)
	path.subpaths = subpaths

	return path
}
