package main

import (
	. "./lib"
	. "./lib/effect"
	. "./lib/utils"
	"fmt"
)

func main() {
	img := Load("../resources/test.jpg")

	ke := Parallel(100, 1000)
	p := ke.Effect().ColorMapper([]MappingRule{Grayscale}).Apply(img)

	//path := NewSegment(Pt2D(-100, -100), Pt2D(100, 100))
	//line := NewLine(path, NewSprayRenderer(10, color.White))
	//line.TransferTo(ke)
	//line.

	//lens := NewFishEyeLens(Pt2D(300, 300), 100, 30)
	//d := NewDistortion(Extend, lens)
	////d.TransferTo(ke)
	//pd := d.Apply(img)

	if err := Save(p.Result(), "../resources/result", "jpeg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
