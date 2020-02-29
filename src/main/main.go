package main

import (
	lib "./lib"
	effect "./lib/effect"
	strategy "./lib/effect/strategy"
	wrap "./lib/effect/wrap"
	"image"

	"fmt"
)

func main() {
	img := lib.Load("../resources/test.jpg")

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

	o := effect.Overlay{Stamp: img, Origin: image.Pt(300, 300)}
	o.TransferTo(&ke)
	p2 := o.Apply(p.Result())

	if err := lib.Save(p2.Result(), "../resources/result", "jpg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
