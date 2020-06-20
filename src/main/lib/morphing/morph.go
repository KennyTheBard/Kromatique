package morphing

import (
	"image"
	"image/draw"

	"../core"
	"../utils"
)

func Morph(engine core.Engine, src, dst image.Image, srcPoints, dstPoints []Vertex, numSteps int) []image.Image {
	if len(srcPoints) != len(dstPoints) {
		panic("Unequal number of points for images")
	}

	src2dst, dst2src := make(map[Vertex]Vertex), make(map[Vertex]Vertex)
	for idx := range srcPoints {
		src2dst[srcPoints[idx]] = dstPoints[idx]
		dst2src[dstPoints[idx]] = srcPoints[idx]
	}

	srcMesh, dstMesh := NewMeshForImage(src), NewMeshForImage(dst)
	BowyerWatson(srcMesh, srcPoints)
	BowyerWatson(dstMesh, dstPoints)

	blend := func(fg, bg image.Image, alpha float64) image.Image {
		for y := fg.Bounds().Min.Y; y < fg.Bounds().Max.Y; y++ {
			for x := fg.Bounds().Min.X; x < fg.Bounds().Max.X; x++ {
				bg.(draw.Image).Set(x, y, utils.PixelLERP(fg.At(x, y), bg.At(x, y), alpha))
			}
		}

		return bg
	}

	stepSize := 1.0 / float64(numSteps)
	ts := make([]float64, 0)
	for t, i := 0.0, 0; i < numSteps; t, i = t+stepSize, i+1 {
		ts = append(ts, t)
	}
	ts = append(ts, 1.0)

	fgPromises := make([]*core.Promise, len(ts))
	bgPromises := make([]*core.Promise, len(ts))
	for i, t := range ts {
		fgPromises[i] = NewMeshDeformation(engine, srcMesh, src2dst, t).Deform(src)
		bgPromises[i] = NewMeshDeformation(engine, dstMesh, dst2src, 1-t).Deform(dst)
	}

	images := make([]image.Image, 0)
	for i, t := range ts {
		images = append(images, blend(fgPromises[i].Result(), bgPromises[i].Result(), t))
	}
	return images
}
