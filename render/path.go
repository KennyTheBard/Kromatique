package render

import (
	"image"
	"image/draw"
	"math"

	"github.com/kennythebard/kromatique/geometry"
	"github.com/kennythebard/kromatique/utils"
)

const maxDistBetween2Points = 1

func skeletonSubpath(path geometry.Path, start, end float64, startPoint, endPoint geometry.Point2D) []geometry.Point2D {
	mid := (start + end) / 2
	midPoint := path.GetPoint(mid)
	var returnPoints []geometry.Point2D

	//fmt.Println(mid, midPoint, start, startPoint, midPoint.Dist(startPoint))
	if midPoint.Dist(startPoint) > maxDistBetween2Points {
		returnPoints = append(returnPoints, skeletonSubpath(path, start, mid, startPoint, midPoint)...)
	} else {
		returnPoints = append(returnPoints, startPoint)
	}

	//fmt.Println(mid, midPoint, end, endPoint, midPoint.Dist(startPoint))
	if midPoint.Dist(endPoint) > maxDistBetween2Points {
		returnPoints = append(returnPoints, skeletonSubpath(path, mid, end, midPoint, endPoint)...)
	} else {
		returnPoints = append(returnPoints, midPoint)
	}

	return returnPoints
}

// PathRender renders a Path with a given Painter and width, using a divide-et impera
// in order to obtain close points
func PathRender(path geometry.Path, paint Painter, width float64) image.Image {
	points := skeletonSubpath(path, 0, 1, path.GetPoint(0), path.GetPoint(1))

	start, end := geometry.Pt2D(0, 0), geometry.Pt2D(0, 0)
	for _, p := range points {
		if p.X < start.X {
			start.X = p.X
		}

		if p.Y < start.Y {
			start.Y = p.Y
		}

		if p.X > end.X {
			end.X = p.X
		}

		if p.Y > end.Y {
			end.Y = p.Y
		}
	}
	ret := utils.CreateRGBA(image.Rect(int(math.Floor(start.X-width-1)), int(math.Floor(start.Y-width-1)), int(math.Ceil(end.X+width+1)), int(math.Ceil(end.Y+width+1))))

	for _, p := range points {
		for y := int(math.Floor(p.Y - width)); y <= int(math.Ceil(p.Y+width)); y++ {
			for x := int(math.Floor(p.X - width)); x <= int(math.Ceil(p.X+width)); x++ {
				op := geometry.Pt2D(float64(x), float64(y))
				if op.DistSq(p) <= width*width {
					ret.(draw.Image).Set(x, y, paint(1, x, y))
				}
			}
		}
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
