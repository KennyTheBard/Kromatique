package geometry

import (
	"image"
	"math"
)

// Shape is an interface that encapsulates any kind of 2D shape
type Shape interface {
	IObject2D
	// Contains returns if the shape contains the given 2D point
	Contains(Point2D) bool
	// MBR = Minimum Bounding Rectangle returns the smallest rectangle
	// that contains the entire shape; should be used to minimize the
	// Inside method calls
	MBR() image.Rectangle
}

// Circle encapsulates a simple circle shape defined by
// a center and a radius
type Circle struct {
	Object2D
	center Point2D
	radius float64
}

func (shape Circle) Contains(p Point2D) bool {
	ip := shape.Inverse().Apply(p)
	return shape.center.Dist(ip) <= shape.radius
}

func (shape Circle) MBR() image.Rectangle {
	min := shape.Model().Apply(Pt2D(shape.center.X-shape.radius, shape.center.Y-shape.radius))
	max := shape.Model().Apply(Pt2D(shape.center.X+shape.radius, shape.center.Y+shape.radius))
	return image.Rect(
		int(math.Floor(min.X)),
		int(math.Floor(min.Y)),
		int(math.Ceil(max.X)),
		int(math.Ceil(max.Y)))
}

func NewCircle(center Point2D, radius float64) *Circle {
	shape := new(Circle)
	InitObject(&shape.Object2D)
	shape.center = center
	shape.radius = radius

	return shape
}

// Rectangle encapsulates a simple rectangle defined by
// an encapsulated image.Rectangle object
type Rectangle struct {
	Object2D
	start, end Point2D
}

func (shape Rectangle) Contains(p Point2D) bool {
	ip := shape.Inverse().Apply(p)
	return ip.X >= shape.start.X && ip.Y >= shape.end.Y && ip.X <= shape.end.X && ip.Y <= shape.end.Y
}

func (shape Rectangle) MBR() image.Rectangle {
	points := []Point2D{
		shape.start, shape.end, Pt2D(shape.end.X, shape.start.Y), Pt2D(shape.start.X, shape.end.Y),
	}

	minX, minY, maxX, maxY := points[0].X, points[0].Y, points[0].X, points[0].Y
	for i := 1; i < len(points); i++ {
		if points[i].X < minX {
			minX = points[i].X
		}

		if points[i].Y < minY {
			minY = points[i].Y
		}

		if points[i].X > maxX {
			maxX = points[i].X
		}

		if points[i].Y > maxY {
			maxY = points[i].Y
		}
	}

	min := shape.Model().Apply(Pt2D(minX, minY))
	max := shape.Model().Apply(Pt2D(maxX, maxY))

	return image.Rect(
		int(math.Floor(min.X)),
		int(math.Floor(min.Y)),
		int(math.Ceil(max.X)),
		int(math.Ceil(max.Y)))
}

func NewRectangle(start, end Point2D) *Rectangle {
	shape := new(Rectangle)
	InitObject(&shape.Object2D)
	shape.start = start
	shape.end = end

	return shape
}
