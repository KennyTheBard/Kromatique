package main

import (
	. "./lib/geometry"
	. "./lib/imageio"
	. "./lib/render"
	. "./lib/utils"
	"fmt"
	"image/color"
	"math"
)

func collide(p1, p2, rayTip Point2D) bool {
	A, B := p2.Y-p1.Y, p1.X-p2.X
	C := (p2.X * p1.Y) - (p1.X * p2.Y)
	x := -(B*rayTip.Y + C) / A
	return x <= rayTip.X
}

func main() {
	//img, err := Load("../resources/test.jpg")
	//if err != nil {
	//	panic(err)
	//}

	//ke := Parallel(100, 1000)

	circle := NewPolygon([]Point2D{
		Pt2D(0, 0),
		Pt2D(100, 0),
		Pt2D(0, 101),
		Pt2D(100, 101),
	})
	renderedImage := ShapeRender(circle, MattePainter(color.RGBA{
		R: math.MaxUint8 - 1,
		G: 0,
		B: 0,
		A: math.MaxUint8 - 1,
	}))

	//p := NewBezier4(
	//	Pt2D(0, 0),
	//	Pt2D(0, 15),
	//	Pt2D(0, 45),
	//	Pt2D(60, 60),
	//)
	//p.Translate(10, 10)
	//renderedImage := PathRender(p, MattePainter(color.RGBA{
	//	R: 255,
	//	G: 0,
	//	B: 0,
	//	A: 255,
	//}), 10)

	if err := Save(renderedImage, "../resources/render", "png"); err != nil {
		fmt.Println(err.Error())
	}

	//ke.Stop()
}
