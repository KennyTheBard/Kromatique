package geometry

import (
	"math"
)

// Point2D encapsulates the coordinates for a point in 2D space
type Point2D struct {
	X, Y float64
}

// Translate applies a translation with the given values
func (p *Point2D) Translate(x, y float64) {
	p.X += x
	p.Y += y
}

// Scale applies a scaling with the given values
func (p *Point2D) Scale(x, y float64) {
	p.X *= x
	p.Y *= y
}

// Rotate applies a rotation with the given radians
func (p *Point2D) Rotate(a float64) {
	p.X = p.X*math.Cos(a) - p.Y*math.Sin(a)
	p.Y = p.X*math.Sin(a) + p.Y*math.Cos(a)
}

// Dist returns the euclidean distance between
// the current Point2D and the given Point2D
func (p Point2D) Dist(q Point2D) float64 {
	return math.Sqrt((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y))
}

// DistSq returns the euclidean distance squared between
// the current Point2D and the given Point2D
func (p Point2D) DistSq(q Point2D) float64 {
	return (p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y)
}

// Diff returns the difference between
// the current Point2D and the given Point2D
func (p Point2D) Diff(q Point2D) Point2D {
	return Pt2D(p.X-q.X, p.Y-q.Y)
}

// Pt2D returns a Point2D with the given coordinates
func Pt2D(X, Y float64) Point2D {
	return Point2D{X: X, Y: Y}
}
