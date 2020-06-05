package render

import (
	"../geometry"
	"../utils"
	"image"
	"image/draw"
	"math"
)

const MaxDistBetween2Points = 1

func skeletonSubpath(path geometry.Path, start, end float64, startPoint, endPoint geometry.Point2D) []geometry.Point2D {
	mid := (start + end) / 2
	midPoint := path.GetPoint(mid)
	var returnPoints []geometry.Point2D

	//fmt.Println(mid, midPoint, start, startPoint, midPoint.Dist(startPoint))
	if midPoint.Dist(startPoint) > MaxDistBetween2Points {
		returnPoints = append(returnPoints, skeletonSubpath(path, start, mid, startPoint, midPoint)...)
	} else {
		returnPoints = append(returnPoints, startPoint)
	}

	//fmt.Println(mid, midPoint, end, endPoint, midPoint.Dist(startPoint))
	if midPoint.Dist(endPoint) > MaxDistBetween2Points {
		returnPoints = append(returnPoints, skeletonSubpath(path, mid, end, midPoint, endPoint)...)
	} else {
		returnPoints = append(returnPoints, midPoint)
	}

	return returnPoints
}

func PathRender(path geometry.Path, paint Painter, width float64) image.Image {
	points := skeletonSubpath(path, 0, 1, path.GetPoint(0), path.GetPoint(1))
	ret := utils.CreateRGBA(image.Rect(0, 0, 100, 100))

	for _, p := range points {
		x, y := int(math.Round(p.X)), int(math.Round(p.Y))
		ret.(draw.Image).Set(x, y, paint(x, y))
	}

	//for y := mbr.Min.Y; y < mbr.Max.Y; y++ {
	//	for x := mbr.Min.X; x < mbr.Max.X; x++ {
	//		if shape.Contains(geometry.Pt2D(float64(x), float64(y))) {
	//			ret.(draw.Image).Set(x, y, paint(x, y))
	//		}
	//	}
	//}

	return ret
}
