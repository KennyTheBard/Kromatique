package main

import (
	lib "./lib"
	effect "./lib/effect"
	strategy "./lib/strategy"

	"fmt"
)

func main() {
	img := lib.Load("../resources/test.jpg")

	ke := lib.NewKromEngine(10, 0)
	f := effect.SingleKernel{}
	f.TransferTo(&ke)
	f.Strategy = strategy.Extend
	f.Matrix = [][]float64{
		{-1, -1, -1},
		{-1, 8, -1},
		{-1, -1, -1},
	}
	p := f.Apply(img)

	if err := lib.Save(p.Result(), "../resources/result", "jpg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
