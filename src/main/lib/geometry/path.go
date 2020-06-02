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
	start, end Point2D
}

func (path *Segment) Translate(x, y float64) {
	path.start.Translate(x, y)
	path.end.Translate(x, y)
}

func (path *Segment) Scale(x, y float64) {
	path.start.Scale(x, y)
	path.end.Scale(x, y)
}

func (path *Segment) Rotate(a float64) {
	path.start.Rotate(a)
	path.end.Rotate(a)
}

func (path *Segment) GetPoint(t float64) Point2D {
	return Pt2D(
		(1-t)*path.start.X+t*path.end.X,
		(1-t)*path.start.Y+t*path.end.Y)
}

func NewSegment(start, end Point2D) *Segment {
	path := new(Segment)
	path.start = start
	path.end = end

	return path
}

// Bezier2 implements a linear Bezier curve
type Bezier2 struct {
	p0, p1 Point2D
}

func (path *Bezier2) Translate(x, y float64) {
	path.p0.Translate(x, y)
	path.p1.Translate(x, y)
}

func (path *Bezier2) Scale(x, y float64) {
	path.p0.Scale(x, y)
	path.p1.Scale(x, y)
}

func (path *Bezier2) Rotate(a float64) {
	path.p0.Rotate(a)
	path.p1.Rotate(a)
}

func (path *Bezier2) GetPoint(t float64) Point2D {
	return Pt2D(
		(1-t)*path.p0.X+t*path.p1.X,
		(1-t)*path.p0.Y+t*path.p1.Y)
}

func NewBezier2(p0, p1 Point2D) *Bezier2 {
	path := new(Bezier2)
	path.p0 = p0
	path.p1 = p1

	return path
}

// Bezier3 implements a quadratic Bezier curve
type Bezier3 struct {
	p0, p1, p2 Point2D
}

func (path *Bezier3) Translate(x, y float64) {
	path.p0.Translate(x, y)
	path.p1.Translate(x, y)
	path.p2.Translate(x, y)
}

func (path *Bezier3) Scale(x, y float64) {
	path.p0.Scale(x, y)
	path.p1.Scale(x, y)
	path.p2.Scale(x, y)
}

func (path *Bezier3) Rotate(a float64) {
	path.p0.Rotate(a)
	path.p1.Rotate(a)
	path.p2.Rotate(a)
}

func (path *Bezier3) GetPoint(t float64) Point2D {
	return Pt2D(
		(1-t)*(1-t)*path.p0.X+2*t*(1-t)*path.p1.X+t*t*path.p2.X,
		(1-t)*(1-t)*path.p0.Y+2*t*(1-t)*path.p1.Y+t*t*path.p2.Y)
}

func NewBezier3(p0, p1, p2 Point2D) *Bezier3 {
	path := new(Bezier3)
	path.p0 = p0
	path.p1 = p1
	path.p2 = p2

	return path
}

// Bezier4 implements a cubic Bezier curve
type Bezier4 struct {
	p0, p1, p2, p3 Point2D
}

func (path *Bezier4) Translate(x, y float64) {
	path.p0.Translate(x, y)
	path.p1.Translate(x, y)
	path.p2.Translate(x, y)
	path.p3.Translate(x, y)
}

func (path *Bezier4) Scale(x, y float64) {
	path.p0.Scale(x, y)
	path.p1.Scale(x, y)
	path.p2.Scale(x, y)
	path.p3.Scale(x, y)
}

func (path *Bezier4) Rotate(a float64) {
	path.p0.Rotate(a)
	path.p1.Rotate(a)
	path.p2.Rotate(a)
	path.p3.Rotate(a)
}

func (path *Bezier4) GetPoint(t float64) Point2D {
	return Pt2D(
		(1-t)*(1-t)*(1-t)*path.p0.X+3*t*(1-t)*(1-t)*path.p1.X+3*t*t*(1-t)*path.p2.X+t*t*t*path.p3.X,
		(1-t)*(1-t)*(1-t)*path.p0.Y+3*t*(1-t)*(1-t)*path.p1.Y+3*t*t*(1-t)*path.p2.Y+t*t*t*path.p3.Y)
}

func NewBezier4(p0, p1, p2, p3 Point2D) *Bezier4 {
	path := new(Bezier4)
	path.p0 = p0
	path.p1 = p1
	path.p2 = p2
	path.p3 = p3

	return path
}

// Hermite implements a cubic Hermite curve
type Hermite struct {
	p0, p1, m0, m1 Point2D
}

func (path *Hermite) Translate(x, y float64) {
	path.p0.Translate(x, y)
	path.p1.Translate(x, y)
	path.m0.Translate(x, y)
	path.m1.Translate(x, y)
}

func (path *Hermite) Scale(x, y float64) {
	path.p0.Scale(x, y)
	path.p1.Scale(x, y)
	path.m0.Scale(x, y)
	path.m1.Scale(x, y)
}

func (path *Hermite) Rotate(a float64) {
	path.p0.Rotate(a)
	path.p1.Rotate(a)
	path.m0.Rotate(a)
	path.m1.Rotate(a)
}

func (path *Hermite) GetPoint(t float64) Point2D {
	return Pt2D(
		(2*t*t*t-3*t*t+1)*path.p0.X+(t*t*t-2*t*t+t)*path.m0.X+(-2*t*t*t+3*t*t)*path.p1.X+(t*t*t-t*t)*path.m1.X,
		(2*t*t*t-3*t*t+1)*path.p0.Y+(t*t*t-2*t*t+t)*path.m0.Y+(-2*t*t*t+3*t*t)*path.p1.Y+(t*t*t-t*t)*path.m1.Y)
}

func NewHermite(p0, p1, m0, m1 Point2D) *Hermite {
	path := new(Hermite)
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
