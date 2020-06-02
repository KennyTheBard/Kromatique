package render

//import (
//	"image"
//	"image/draw"
//
//	"../geometry"
//)
//
//func PathRender(path geometry.Path, paint Painter) image.Image {
//	points := make([]geometry.Point2D, 2)
//	ret := utils.CreateRGBA(image.Rect(0, 0, mbr.Max.X, mbr.Min.Y))
//
//	for y := mbr.Min.Y; y < mbr.Max.Y; y++ {
//		for x := mbr.Min.X; x < mbr.Max.X; x++ {
//			if shape.Contains(geometry.Pt2D(float64(x), float64(y))) {
//				ret.(draw.Image).Set(x, y, paint(x, y))
//			}
//		}
//	}
//
//	return ret
//}
