package main

import (
	. "./lib"
	. "./lib/histogram"
	. "./lib/utils"
	"fmt"
	"math"
)

func main() {
	img := Load("../resources/test_eq.jpg")

	ke := Parallel(100, 1000)
	h := NewHistogram(LightnessEvaluation, 0, math.MaxUint8)
	h.Scan(img)
	fmt.Println(len(h.Values()), h.Values())

	newImg := Equalize(h, UniformColorShift)
	h.Scan(newImg)
	fmt.Println(len(h.Values()), h.Values())

	if err := Save(newImg, "../resources/result", "png"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
