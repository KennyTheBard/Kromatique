package main

import (
	. "./lib"
	. "./lib/geometry"
	. "./lib/histogram"
	. "./lib/render"
	. "./lib/utils"
	"fmt"
	"image"
	"image/color"
	"math"
)

func main() {
	img := Load("../resources/test_eq.jpg")

	ke := Parallel(100, 1000)
	h := NewHistogram(LightnessEvaluation, math.MaxUint8+1)
	h.Scan(img)

	newImg := Equalize(img, h, UniformColorShift)
	h.Scan(newImg)

	if err := Save(newImg, "../resources/result", "png"); err != nil {
		fmt.Println(err.Error())
	}

	bg := Load("../resources/bg.jpg")
	fg := Load("../resources/rgb.jpg")

	p := ke.Blend().Divide()(bg, fg, image.Pt(0, 0))
	if err := Save(p.Result(), "../resources/blend", "png"); err != nil {
		fmt.Println(err.Error())
	}

	circle := NewCircle(Pt2D(100, 100), 35)
	circle.Translate(-40, -40)
	fmt.Println(circle.MBR())
	renderedImage := ShapeRender(circle, MattePainter(color.RGBA{
		R: math.MaxUint8 - 1,
		G: 0,
		B: 0,
		A: math.MaxUint8 - 1,
	}))
	if err := Save(renderedImage, "../resources/render", "png"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
