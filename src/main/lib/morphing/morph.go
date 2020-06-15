package morphing

import "image"

func Morph(src, dst image.Image, srcPoints, dstPoints []Vertex) {
	if len(srcPoints) != len(dstPoints) {
		panic("Unequal number of points for images")
	}

	src2dst, dst2src := make(map[Vertex]Vertex), make(map[Vertex]Vertex)
	for idx := range srcPoints {
		src2dst[srcPoints[idx]] = dstPoints[idx]
		dst2src[dstPoints[idx]] = srcPoints[idx]
	}

	srcMesh, dstMesh := NewMesh(src), NewMesh(dst)
	BowyerWatson(srcMesh, srcPoints)
	BowyerWatson(dstMesh, dstPoints)

	NewMeshDeformation(srcMesh, dst, src2dst, 0)
	NewMeshDeformation(dstMesh, src, src2dst, 0)
}
