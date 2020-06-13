package morphing

import (
	"../geometry"
)

func BowyerWatson(points []geometry.Point2D, start, end geometry.Point2D) *Mesh {
	triangulation := NewMesh(start, end)
	triangles := triangulation.Triangles

	for _, p := range points {
		badTriangles := make([]int, 0)
		for idx, tri := range triangles {
			if p.DistSq(tri.circle.center) <= tri.circle.radius*tri.circle.radius {
				badTriangles = append(badTriangles, idx)
			}
		}

		type edge struct {
			start, end geometry.Point2D
		}

		polygon := make([]edge, 0)
		for idx1, i := range badTriangles {
			tmp1 := triangles[i]
			check1, check2, check3 := true, true, true
			for idx2, j := range badTriangles {
				if idx1 == idx2 {
					continue
				}

				tmp2 := triangles[j]
				check1 = check1 && !(tmp2.HasPoint(tmp1.points[0]) && tmp2.HasPoint(tmp1.points[1]))
				check2 = check2 && !(tmp2.HasPoint(tmp1.points[1]) && tmp2.HasPoint(tmp1.points[2]))
				check3 = check3 && !(tmp2.HasPoint(tmp1.points[2]) && tmp2.HasPoint(tmp1.points[0]))
			}

			if check1 {
				polygon = append(polygon, edge{
					start: tmp1.points[0],
					end:   tmp1.points[1],
				})
			}
			if check2 {
				polygon = append(polygon, edge{
					start: tmp1.points[1],
					end:   tmp1.points[2],
				})
			}
			if check3 {
				polygon = append(polygon, edge{
					start: tmp1.points[2],
					end:   tmp1.points[0],
				})
			}

		}

		temp := make([]*Triangle, 0)
	trianglesLoop:
		for idx, tri := range triangles {
			for _, badIdx := range badTriangles {
				if idx == badIdx {
					continue trianglesLoop
				}
			}

			temp = append(temp, tri)
		}
		triangles = temp

		for _, e := range polygon {
			triangles = append(triangles, NewTriangle(e.start, e.end, p))
		}
	}

	triangulation.Triangles = triangles
	return triangulation
}
