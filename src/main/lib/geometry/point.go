package geometry

import (
	"math"
)

type Point2D struct {
	X, Y float64
}

func (p *Point2D) Translate(x, y float64) {
	p.X += x
	p.Y += y
}

func (p *Point2D) Scale(x, y float64) {
	p.X *= x
	p.Y *= y
}

func (p *Point2D) Rotate(a float64) {
	p.X = p.X*math.Cos(a) - p.Y*math.Sin(a)
	p.Y = p.X*math.Sin(a) + p.Y*math.Cos(a)
}

func (p Point2D) Dist(q Point2D) float64 {
	return math.Sqrt((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y))
}

func (p Point2D) DistSq(q Point2D) float64 {
	return (p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y)
}

func (p Point2D) Diff(q Point2D) Point2D {
	return Pt2D(p.X-q.X, p.Y-q.Y)
}

func Pt2D(X, Y float64) Point2D {
	return Point2D{X: X, Y: Y}
}
