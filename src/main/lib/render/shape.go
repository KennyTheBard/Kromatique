package render

import (
	"image"
	"image/draw"

	"../geometry"
	"../utils"
)

func ShapeRender(shape geometry.Shape, paint Painter) image.Image {
	mbr := shape.MBR()
	ret := utils.CreateRGBA(image.Rect(0, 0, mbr.Max.X, mbr.Max.Y))

	for y := mbr.Min.Y; y < mbr.Max.Y; y++ {
		for x := mbr.Min.X; x < mbr.Max.X; x++ {
			if shape.Contains(geometry.Pt2D(float64(x), float64(y))) {
				c := paint(x, y)
				ret.(draw.Image).Set(x, y, c)
			}
		}
	}

	return ret
}
