package morphing

import (
	"../geometry"
)

type SimpleCircle struct {
	center geometry.Point2D
	radius float64
}

type Triangle struct {
	points [3]geometry.Point2D
	circle SimpleCircle
}

func (t *Triangle) HasPoint(p geometry.Point2D) bool {
	for _, tp := range t.points {
		if p.DistSq(tp) < 0.001 {
			return true
		}
	}
	return false
}

func circumscribedCircle(p0, p1, p2 geometry.Point2D) SimpleCircle {
	ax, ay := p1.X-p0.X, p1.Y-p0.Y
	bx, by := p2.X-p0.X, p2.Y-p0.Y

	m := p1.X*p1.X - p0.X*p0.X + p1.Y*p1.Y - p0.Y*p0.Y
	u := p2.X*p2.X - p0.X*p0.X + p2.Y*p2.Y - p0.Y*p0.Y
	s := 1.0 / (2.0 * (ax*by - ay*bx))

	center := geometry.Pt2D(((p2.Y-p0.Y)*m+(p0.Y-p1.Y)*u)*s, ((p0.X-p2.X)*m+(p1.X-p0.X)*u)*s)

	return SimpleCircle{
		center: center,
		radius: center.Dist(p0),
	}
}

func NewTriangle(p0, p1, p2 geometry.Point2D) *Triangle {
	tri := new(Triangle)
	tri.points = [3]geometry.Point2D{p0, p1, p2}
	tri.circle = circumscribedCircle(p0, p1, p2)

	return tri
}

type Mesh struct {
	Triangles []*Triangle
}

func NewMesh(start, end geometry.Point2D) *Mesh {
	mesh := new(Mesh)
	mesh.Triangles = make([]*Triangle, 2)
	mesh.Triangles[0] = NewTriangle(start, geometry.Pt2D(start.X, end.Y), end)
	mesh.Triangles[1] = NewTriangle(start, geometry.Pt2D(end.X, start.Y), end)

	return mesh
}
