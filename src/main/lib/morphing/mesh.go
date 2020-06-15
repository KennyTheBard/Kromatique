package morphing

import (
	"image"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (p Vertex) Dist(q Vertex) float64 {
	return math.Sqrt(p.DistSq(q))
}

func (p Vertex) DistSq(q Vertex) float64 {
	return (p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y)
}

func (p Vertex) Equal(q Vertex) bool {
	return p.DistSq(q) < 0.0001
}

func Vx(X, Y float64) Vertex {
	return Vertex{X: X, Y: Y}
}

type Circle struct {
	center   Vertex
	radiusSq float64
}

type Triangle struct {
	points [3]Vertex
	circle Circle
}

func (t *Triangle) IsVertex(p Vertex) bool {
	for _, tp := range t.points {
		if p.Equal(tp) {
			return true
		}
	}
	return false
}

func (t *Triangle) IsEdge(e Edge) bool {
	for _, te := range t.Edges() {
		if te.Equal(e) {
			return true
		}
	}

	return false
}

func (t *Triangle) HasVertex(v Vertex) bool {
	sign := func(p1, p2, p3 Vertex) float64 {
		return (p1.X-p3.X)*(p2.Y-p3.Y) - (p2.X-p3.X)*(p1.Y-p3.Y)
	}

	d1 := sign(v, t.points[0], t.points[1])
	d2 := sign(v, t.points[1], t.points[2])
	d3 := sign(v, t.points[2], t.points[0])

	hasNeg := (d1 < 0) || (d2 < 0) || (d3 < 0)
	hasPos := (d1 > 0) || (d2 > 0) || (d3 > 0)

	return !(hasNeg && hasPos)
}

func (t *Triangle) Edges() [3]Edge {
	return [3]Edge{
		{
			Start: t.points[0],
			End:   t.points[1],
		},
		{
			Start: t.points[1],
			End:   t.points[2],
		},
		{
			Start: t.points[2],
			End:   t.points[0],
		},
	}
}

func circumscribedCircle(p0, p1, p2 Vertex) Circle {
	ax, ay := p1.X-p0.X, p1.Y-p0.Y
	bx, by := p2.X-p0.X, p2.Y-p0.Y

	m := p1.X*p1.X - p0.X*p0.X + p1.Y*p1.Y - p0.Y*p0.Y
	u := p2.X*p2.X - p0.X*p0.X + p2.Y*p2.Y - p0.Y*p0.Y
	s := 1.0 / (2.0 * (ax*by - ay*bx))

	center := Vx(((p2.Y-p0.Y)*m+(p0.Y-p1.Y)*u)*s, ((p0.X-p2.X)*m+(p1.X-p0.X)*u)*s)

	return Circle{
		center:   center,
		radiusSq: center.DistSq(p0),
	}
}

func NewTriangle(p0, p1, p2 Vertex) *Triangle {
	tri := new(Triangle)
	tri.points = [3]Vertex{p0, p1, p2}
	tri.circle = circumscribedCircle(p0, p1, p2)

	return tri
}

type Edge struct {
	Start, End Vertex
}

func (e *Edge) Equal(oe Edge) bool {
	return (e.Start.Equal(oe.Start) && e.End.Equal(oe.End)) || (e.Start.Equal(oe.End) && e.End.Equal(oe.Start))
}

type Mesh struct {
	Triangles []Triangle
	Texture   image.Image
}

func (m *Mesh) Edges() ([]Edge, map[Edge][]Triangle) {
	edges := make([]Edge, 0)
	ownership := make(map[Edge][]Triangle)
	for _, tri := range m.Triangles {
		tmpEdges := tri.Edges()

	edgeLoop:
		for _, tmpEdge := range tmpEdges {
			for _, e := range edges {
				if tmpEdge.Equal(e) {
					ownership[e] = append(ownership[e], tri)
					continue edgeLoop
				}
			}

			edges = append(edges, tmpEdge)
			ownership[tmpEdge] = []Triangle{tri}
		}
	}

	return edges, ownership
}

func (m *Mesh) Vertexes() ([]Vertex, map[Vertex][]Triangle) {
	points := make([]Vertex, 0)
	ownership := make(map[Vertex][]Triangle)
	for _, tri := range m.Triangles {
	vertexLoop:
		for _, tmpVertex := range tri.points {
			for _, p := range points {
				if tmpVertex.Equal(p) {
					ownership[p] = append(ownership[p], tri)
					continue vertexLoop
				}
			}

			points = append(points, tmpVertex)
			ownership[tmpVertex] = []Triangle{tri}
		}
	}

	return points, ownership
}

func NewMesh(texture image.Image) *Mesh {
	start := Vx(float64(texture.Bounds().Min.X), float64(texture.Bounds().Min.Y))
	end := Vx(float64(texture.Bounds().Max.X), float64(texture.Bounds().Max.Y))

	mesh := new(Mesh)
	mesh.Texture = texture
	mesh.Triangles = make([]Triangle, 2)
	mesh.Triangles[0] = *NewTriangle(start, Vx(start.X, end.Y), end)
	mesh.Triangles[1] = *NewTriangle(start, Vx(end.X, start.Y), end)

	return mesh
}
