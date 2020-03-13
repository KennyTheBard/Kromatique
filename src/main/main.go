package main

import (
	"fmt"
	"image"
	"image/color"

	. "./lib"
	. "./lib/analysis"
	. "./lib/effect/filter"
	. "./lib/effect/normalization"
	. "./lib/effect/overlay"
	. "./lib/effect/scale"
	. "./lib/geometry"
	. "./lib/position"
	. "./lib/render"
	. "./lib/utils"
)

func main() {
	img := Load("../resources/test.jpg")

	ke := NewKromEngine(10, 0)
	f := MultiKernel{}
	f.TransferTo(ke)
	f.EdgeHandling = Extend
	f.ResultMerging = SobelGradient
	f.Kernels = []Matrix{
		{
			{1, 0, -1},
			{2, 0, -2},
			{1, 0, -1},
		},
		{
			{1, 2, 1},
			{0, 0, 0},
			{-1, -2, -1},
		},
	}
	p1 := f.Apply(img)

	s := NewScale(
		NewFixedScaleFactor(ScaleFactor{X: 0.71, Y: 0.71}),
		CornerPixelsSampling)
	s.TransferTo(ke)
	p2 := s.Apply(p1.Result())
	tmp := p2.Result()

	data := NewAnalyzerRunner([]Analyze{func(img image.Image, x int, y int, m map[string]interface{}) {
		val, _, _, _ := color.Gray16Model.Convert(img.At(x, y)).RGBA()

		if t, ok := m["min"]; ok {
			if val < t.(uint32) {
				m["min"] = val
			}
		} else {
			m["min"] = val
		}

		if t, ok := m["max"]; ok {
			if val > t.(uint32) {
				m["max"] = val
			}
		} else {
			m["max"] = val
		}
	}}).Run(tmp)

	n := NewNormalization(
		NewColorInterval(data["min"].(uint32), data["max"].(uint32)),
		NewColorInterval(MaxUint16/4, MaxUint16/4*3))
	n.TransferTo(ke)
	//p3 := n.Apply(tmp)
	//
	//path := NewSegment(NewPoint2D(100, 100), NewPoint2D(300, 300))
	//line := NewLine(path, NewSprayRenderer(10, color.RGBA{
	//	R: 255,
	//	G: 0,
	//	B: 0,
	//	A: 255,
	//}))
	//line.TransferTo(ke)
	//p4 := line.Render(image.Rect(0, 0, 600, 600))
	//
	//o := NewOverlay(p4.Result(), NewFixedPosition(100, 100), 1.00)
	//o.TransferTo(ke)
	//p5 := o.Apply(p3.Result())
	//
	//if err := Save(p5.Result(), "../resources/result", "png"); err != nil {
	//	fmt.Println(err.Error())
	//}

	test(ke)

	ke.Stop()
}

func test(ke *KromEngine) {

	//
	//for i := 0.0; i <= 1.0; i += 0.1 {
	//	fmt.Println(PixelLERP(
	//		color.RGBA64{65535, 0, 0, 0},
	//		color.RGBA64{65535, 0, 0, 65535},
	//		i))
	//}

	path := NewSegment(NewPoint2D(15, 15), NewPoint2D(100, 100))
	line := NewLine(path, NewSprayRenderer(10, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}))
	line.TransferTo(ke)
	p := line.Render(image.Rect(0, 0, 120, 120))
	img := p.Result()

	o := NewOverlay(CreateBackground(img.Bounds(), color.Black), NewFixedPosition(0, 0), 1.00)
	o.TransferTo(ke)
	p2 := o.Apply(img)

	if err := Save(p2.Result(), "../resources/result", "jpeg"); err != nil {
		fmt.Println(err.Error())
	}
}
