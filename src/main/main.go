package main

import (
	lib "./lib"
	analysis "./lib/analysis"
	filter "./lib/effect/filter"
	scale "./lib/effect/scale"
	utils "./lib/utils"
	"fmt"
	"image"
	"image/color"
	"math"
)

func main() {
	img := utils.Load("../resources/test.jpg")

	ke := lib.NewKromEngine(10, 0)
	f := filter.MultiKernel{}
	f.TransferTo(ke)
	f.EdgeHandling = filter.Extend
	f.ResultMerging = filter.SobelGradient
	f.Kernels = []utils.Matrix{
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

	s := scale.NewScale(
		scale.NewFixedScaleFactor(scale.ScaleFactor{X: 0.71, Y: 0.71}),
		scale.CornerPixelsSampling)
	s.TransferTo(ke)
	p2 := s.Apply(p1.Result())
	res := p2.Result()

	data := analysis.NewAnalyzerRunner([]analysis.Analyze{func(img image.Image, x int, y int, m map[string]interface{}) {
		var size float64
		if s, ok := m["size"]; ok {
			size = s.(float64)
		} else {
			size = float64(img.Bounds().Dx() * img.Bounds().Dy())
			m["size"] = size
		}

		val, _, _, _ := color.Gray16Model.Convert(img.At(x, y)).RGBA()

		if t, ok := m["total"]; ok {
			m["total"] = t.(float64) + float64(val)/size
		} else {
			m["total"] = float64(val) / size
		}
	}}).Run(res)

	gray := utils.ClampUint16(math.Round(data["total"].(float64)))
	utils.CreateBackground(res.Bounds(), color.Gray16{Y: uint16(gray)})
	fmt.Println(data["total"])

	if err := utils.Save(utils.CreateBackground(res.Bounds(), color.Gray16{Y: uint16(gray)}),
		"../resources/result", "jpg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
