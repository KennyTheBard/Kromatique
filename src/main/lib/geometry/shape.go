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

func (shape *Circle) Contains(p Point2D) bool {
	ip := shape.Inverse().Apply(p)
	return shape.center.Dist(ip) <= shape.radius
}

func (shape *Circle) MBR() image.Rectangle {
	min := shape.Model().Apply(Pt2D(shape.center.X-shape.radius, shape.center.Y-shape.radius))
	max := shape.Model().Apply(Pt2D(shape.center.X+shape.radius, shape.center.Y+shape.radius))
	return image.Rect(
		int(math.Floor(min.X)),
		int(math.Floor(min.Y)),
		int(math.Ceil(max.X)),
		int(math.Ceil(max.Y))).Canon()
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

func (shape *Rectangle) Contains(p Point2D) bool {
	ip := shape.Inverse().Apply(p)
	return ip.X >= shape.start.X && ip.Y >= shape.end.Y && ip.X <= shape.end.X && ip.Y <= shape.end.Y
}

func (shape *Rectangle) MBR() image.Rectangle {
	start := shape.Model().Apply(Pt2D(shape.start.X, shape.start.Y))
	end := shape.Model().Apply(Pt2D(shape.end.X, shape.end.Y))

	return image.Rect(
		int(math.Floor(start.X)),
		int(math.Floor(start.Y)),
		int(math.Ceil(end.X)),
		int(math.Ceil(end.Y))).Canon()
}

func NewRectangle(start, end Point2D) *Rectangle {
	shape := new(Rectangle)
	InitObject(&shape.Object2D)
	shape.start = Pt2D(0, 0)
	shape.end = Pt2D(0, 0)
	if end.X < start.X {
		shape.start.Translate(end.X, 0)
		shape.end.Translate(start.X, 0)
	}
	if end.Y < start.Y {
		shape.start.Translate(0, end.Y)
		shape.end.Translate(0, start.Y)
	}

	return shape
}

// Polygon encapsulates a polygon defined by a list of points,
// where the last is implicitly bound to the first
type Polygon struct {
	Object2D
	points []Point2D
}

func collide(p1, p2, rayTip Point2D) (Point2D, bool) {
	A, B := p2.Y-p1.Y, p1.X-p2.X
	C := (p2.X * p1.Y) - (p1.X * p2.Y)
	x := -(B*rayTip.Y + C) / A
	return Pt2D(x, rayTip.Y), x <= rayTip.X
}

func (shape *Polygon) Contains(p Point2D) bool {
	if len(shape.points) <= 2 {
		return false
	}

	p = shape.Inverse().Apply(p)

	// count collisions of the ray with polygon edges
	rayCollisionPoints := make(map[int]Point2D, 0)
	for i := 0; i < len(shape.points); i++ {
		var prev, curr Point2D
		curr = shape.points[i]
		if i == 0 {
			prev = shape.points[len(shape.points)-1]
		} else {
			prev = shape.points[i-1]
		}

		// check if the edge is parallel with OX (should ignore)
		if prev.Y == curr.Y {
			continue
		}

		// check if the segment intersects the horizontal line
		if !(prev.Y >= p.Y && curr.Y <= p.Y) && !(prev.Y <= p.Y && curr.Y >= p.Y) {
			continue
		}

		// if both points are after ray's tip they cannot collide
		if curr.X > p.X && prev.X > p.X {
			continue
		}

		// check collision for segments with the ray tip inside their MBR
		if collisionPoint, ok := collide(prev, curr, p); ok {
			rayCollisionPoints[i] = collisionPoint
		}
	}

	// count for doubled collisions around definition points
	rayCollisionCount := 0
	for key, currCollisionPoint := range rayCollisionPoints {
		var prevKey, prevPrevKey int
		if key == 0 {
			prevKey = len(shape.points) - 1
			prevPrevKey = len(shape.points) - 2
		} else if key == 1 {
			prevKey = 0
			prevPrevKey = len(shape.points) - 1
		} else {
			prevKey = key - 1
			prevPrevKey = key - 2
		}
		curr := shape.points[key]
		prev := shape.points[prevKey]
		prevPrev := shape.points[prevPrevKey]

		// if there are 2 collisions for 2 edges that are consecutive
		// and each on a different side of the collision ray,
		// only one should be taken into account
		if prevCollisionPoint, ok := rayCollisionPoints[prevKey]; ok {
			if currCollisionPoint.Dist(prevCollisionPoint) < 0.001 {
				if math.Signbit(curr.Y-prev.Y) == math.Signbit(prev.Y-prevPrev.Y) {
					continue
				}
			}
		}

		rayCollisionCount += 1
	}

	return rayCollisionCount%2 == 1
}

func (shape *Polygon) MBR() image.Rectangle {
	if len(shape.points) <= 2 {
		return image.Rectangle{}
	}

	minX, minY, maxX, maxY := shape.points[0].X, shape.points[0].Y, shape.points[0].X, shape.points[0].Y
	for i := 1; i < len(shape.points); i++ {
		p := shape.points[i]

		if p.X < minX {
			minX = p.X
		}

		if p.Y < minY {
			minY = p.Y
		}

		if p.X > maxX {
			maxX = p.X
		}

		if p.Y > maxY {
			maxY = p.Y
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

func (shape *Polygon) AddPoint(p Point2D) {
	shape.points = append(shape.points, p)
}

func NewPolygon(points []Point2D) *Polygon {
	shape := new(Polygon)
	InitObject(&shape.Object2D)
	shape.points = points

	return shape
}
