package main

import (
	. "./lib"
	. "./lib/crop"
	. "./lib/position"
	. "./lib/utils"
	"fmt"
)

func main() {
	img, err := Load("../resources/test.jpg")
	if err != nil {
		panic(err)
	}

	ke := Parallel(100, 1000)

	//circle := NewCircle(Pt2D(100, 100), 35)
	//circle.Translate(-40, -40).Scale(1.2, 0.5)
	//fmt.Println(circle.MBR())
	//renderedImage := ShapeRender(circle, MattePainter(color.RGBA{
	//	R: math.MaxUint8 - 1,
	//	G: 0,
	//	B: 0,
	//	A: math.MaxUint8 - 1,
	//}))

	res := Copy(img, Pos(Fixed(200), Fixed(200)), Pos(Fixed(300), Fixed(200).Anchor(End)))

	fmt.Println(res.Bounds())

	if err := Save(res, "../resources/result", "png"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
