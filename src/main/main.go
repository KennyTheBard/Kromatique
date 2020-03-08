package main

import (
	"fmt"
	"image"
	"image/color"

	. "./lib"
	. "./lib/analysis"
	. "./lib/effect/filter"
	. "./lib/effect/normalization"
	. "./lib/effect/scale"
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
	res := p2.Result()

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
	}}).Run(res)

	fmt.Println(data)

	n := NewNormalization(
		NewColorInterval(data["min"].(uint32), data["max"].(uint32)),
		NewColorInterval(MaxUint16/4, MaxUint16/4*3))
	n.TransferTo(ke)
	p3 := n.Apply(res)

	if err := Save(p3.Result(), "../resources/result", "jpg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
