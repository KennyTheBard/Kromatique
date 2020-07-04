package render

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"sort"

	"github.com/kennythebard/kromatique/core"
	"github.com/kennythebard/kromatique/geometry"
	"github.com/kennythebard/kromatique/utils"
)

func ShapeRender(shape geometry.Shape, paint Painter) image.Image {
	switch t := shape.(type) {
	case *geometry.Circle:
		return CircleRender(t, paint)
	case *geometry.Polygon:
		return PolygonRender(t, paint)
	default:
		return image.Black
	}
}

func CircleRender(shape *geometry.Circle, paint Painter) image.Image {
	mbr := shape.MBR()
	ret := utils.CreateRGBA(mbr)

	def := shape.Definition()
	center := def[0].(geometry.Point2D)
	radius := def[1].(float64)

	for y := mbr.Min.Y; y <= mbr.Max.Y; y++ {
		for x := mbr.Min.X; x <= mbr.Max.X; x++ {
			p := geometry.Pt2D(float64(x)+0.5, float64(y)+0.5)
			ip := shape.Inverse().Apply(p)
			dist := ip.Dist(center)

			if dist > radius+0.5 {
				continue
			}

			var c color.Color
			if dist <= radius-0.5 {
				c = paint(1, x, y)
			} else {
				c = paint(1-(dist-(radius-0.5)), x, y)
			}

			ret.(draw.Image).Set(x, y, c)
		}
	}

	return ret
}

func PolygonRender(shape *geometry.Polygon, paint Painter) image.Image {
	mbr := shape.MBR()
	ret := utils.CreateRGBA(mbr)

	// cast shape definition
	def := shape.Definition()
	points := def[0].([]geometry.Point2D)

	// early escape
	if len(points) <= 2 {
		return ret
	}

	// define a support structure to simplify working with edges
	type edge struct {
		start, end *geometry.Point2D
	}
	edgeConstructor := func(start, end *geometry.Point2D) edge {
		return edge{
			start: start,
			end:   end,
		}
	}

	// create a list of all the edges of the  polygon
	edges := make([]edge, len(points))
	for i := 0; i < len(points); i++ {
		var prev, curr int
		curr = i
		if i == 0 {
			prev = len(points) - 1
		} else {
			prev = i - 1
		}

		edges[i] = edgeConstructor(&points[prev], &points[curr])
	}

	// define helper function
	collide := func(e edge, y float64) geometry.Point2D {
		A, B := e.end.Y-e.start.Y, e.start.X-e.end.X
		C := (e.end.X * e.start.Y) - (e.start.X * e.end.Y)
		x := -(B*y + C) / A
		return geometry.Pt2D(x, y)
	}

	// iterate through each line and render contained pixels
	core.Parallelize(mbr.Dy(), func(y int) {
		Y := float64(y) + 0.5

		collisionMap := make(map[int]geometry.Point2D)

		// find edges that have a collision point with the current line
		for idx, e := range edges {
			// horizontal lines should be ignored
			if e.start.Y == e.end.Y {
				continue
			}

			// if both defining points of an edge are on the same side
			// of the horizontal line, there cannot be any collision
			if !(e.start.Y >= Y && e.end.Y <= Y) && !(e.start.Y <= Y && e.end.Y >= Y) {
				continue
			}

			collisionPoint := collide(e, Y)
			collisionMap[idx] = collisionPoint
		}

		// remove duplicate collision points for adjacent edges
		// that are on opposite sides of the current line
		collisionPoints := make([]geometry.Point2D, 0)
		for idx, collisionPoint := range collisionMap {
			var prevIdx int
			if idx == 0 {
				prevIdx = len(edges) - 1
			} else {
				prevIdx = idx - 1
			}

			// check if the collision point should be skipped
			// if there is another for the previous adjacent edge
			prevCollisionPoint, present := collisionMap[prevIdx]
			if present {
				prev := edges[prevIdx]
				curr := edges[idx]
				if math.Abs(prevCollisionPoint.X-collisionPoint.X) < 0.001 {
					if math.Signbit(curr.start.Y-prev.end.Y) == math.Signbit(prev.start.Y-prev.end.Y) {
						continue
					}
				}
			}

			collisionPoints = append(collisionPoints, collisionPoint)
		}

		// sort the points in the order that an horizontal
		// ray going from left to right would pass them
		sort.Slice(collisionPoints[:], func(i, j int) bool {
			return collisionPoints[i].X < collisionPoints[j].X
		})

		// render each point of the line
		next := 0
		for x := mbr.Min.X; x <= mbr.Max.X; x++ {
			X := float64(x)

			delta := 1.0
			if next%2 == 0 {
				delta *= -1
			}

			for next < len(collisionPoints) && X > collisionPoints[next].X+0.5*delta {
				next += 1
			}

			if next > 0 {
				delta := 1.0
				if next%2 == 1 {
					delta *= -1
				}

				prev := next - 1
				var alpha float64
				if next == len(collisionPoints) {
					alpha = delta * (X - collisionPoints[prev].X)
				} else {
					alpha = delta * math.Min(X-collisionPoints[prev].X, collisionPoints[next].X-X)
				}

				if alpha < 0 {
					ret.(draw.Image).Set(x, y, paint(1, x, y))
				} else if alpha <= 0.5 {
					alpha = 1 - (alpha)
					ret.(draw.Image).Set(x, y, paint(alpha, x, y))
				}
			}
		}
	})

	return ret
}
