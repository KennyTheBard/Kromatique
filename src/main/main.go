package main

import (
	lib "./lib"
	effect "./lib/effect"
	strategy "./lib/effect/strategy"
	wrap "./lib/effect/wrap"
	pos "./lib/position"
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
	p := f.Apply(img)

	o := effect.Overlay{
		Stamp: img,
		Origin: pos.NewRelativePosition(
			pos.NewRelativeAxialPosition(100, pos.Center),
			pos.NewRelativeAxialPosition(0, pos.Center)),
		Opacity: 0.9}
	o.TransferTo(&ke)
	p2 := o.Apply(p.Result())

	n := effect.Negative{}
	n.TransferTo(&ke)
	p3 := n.Apply(p2.Result())

	if err := utils.Save(p3.Result(), "../resources/result", "jpg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
