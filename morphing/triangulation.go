package morphing

func BowyerWatson(mesh *Mesh, points []Vertex) {
	triangles := mesh.Triangles

	for _, p := range points {
		badTriangles := make([]int, 0)
		for idx, tri := range triangles {
			if p.DistSq(tri.circle.center) <= tri.circle.radiusSq {
				badTriangles = append(badTriangles, idx)
			}
		}

		polygon := make([]Edge, 0)
		for idx1, i := range badTriangles {
		mainEdgeLoop:
			for _, badEdge := range triangles[i].Edges() {
				for idx2, j := range badTriangles {
					if idx1 == idx2 {
						continue
					}

					for _, otherBadEdge := range triangles[j].Edges() {
						if badEdge.Equal(otherBadEdge) {
							continue mainEdgeLoop
						}
					}
				}

				polygon = append(polygon, badEdge)
			}
		}

		temp := make([]Triangle, 0)
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
			triangles = append(triangles, *NewTriangle(e.Start, e.End, p))
		}
	}

	mesh.Triangles = triangles
}
