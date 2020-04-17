package main

import (
	. "./lib"
	. "./lib/effect/distortion"
	. "./lib/effect/filter"
	. "./lib/utils"
	"fmt"
)

func main() {
	img := Load("../resources/test.jpg")

	ke := NewKromEngine(10, 0)

	//cmr := NewColorMapperRunner()
	//cmr.TransferTo(ke)
	//cmr.Add(ColorMapperFactory(
	//	func(col color.Color) bool {
	//		r, g, b, _ := col.RGBA()
	//		return r+g+b < math.MaxUint16*3/2
	//	},
	//	func(col color.Color) color.Color {
	//		_, _, _, a := col.RGBA()
	//		return color.RGBA64{
	//			R: uint16(0),
	//			G: uint16(0),
	//			B: uint16(0),
	//			A: uint16(a),
	//		}
	//	}))
	//
	//cmr.Add(ColorMapperFactory(
	//	func(col color.Color) bool {
	//		r, g, b, _ := col.RGBA()
	//		return r+g+b > math.MaxUint16*3/2
	//	},
	//	func(col color.Color) color.Color {
	//		_, _, _, a := col.RGBA()
	//		return color.RGBA64{
	//			R: math.MaxUint16,
	//			G: math.MaxUint16,
	//			B: math.MaxUint16,
	//			A: uint16(a),
	//		}
	//	}))
	//p := cmr.Apply(img)

	lens := NewFishEyeLens(Pt2D(300, 300), 100, 30)
	d := NewDistortion(Extend, lens)
	d.TransferTo(ke)
	pd := d.Apply(img)

	if err := Save(pd.Result(), "../resources/result", "jpeg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
