package main

import (
	lib "./lib"
	effect "./lib/effect"
	strategy "./lib/effect/strategy"
	wrap "./lib/effect/wrap"
	utils "./lib/utils"
	"fmt"
)

func main() {
	img := utils.Load("../resources/test.jpg")

	ke := lib.NewKromEngine(10, 0)
	f := effect.MultiKernel{}
	f.TransferTo(&ke)
	f.EdgeHandling = strategy.Extend
	f.ResultMerging = strategy.SobelGradient
	f.Kernels = []wrap.Matrix{
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

	n := effect.NewScale(
		strategy.NewFixedScaleFactor(strategy.ScaleFactor{X: 0.71, Y: 0.71}),
		strategy.CornerPixelsSampling)
	n.TransferTo(&ke)

	if err := utils.Save(n.Apply(f.Apply(img).Result()).Result(), "../resources/result", "jpg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
