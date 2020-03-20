package main

import (
	"fmt"
	"image/color"
	"math"

	. "./lib"
	. "./lib/effect/difference"
	. "./lib/effect/distortion"
	"./lib/effect/filter"
	. "./lib/effect/mapper"
	. "./lib/effect/mirror"
	. "./lib/utils"
)

func main() {
	img := Load("../resources/test.jpg")

	ke := NewKromEngine(10, 0)

	cmr := NewColorMapperRunner()
	cmr.TransferTo(ke)
	//cmr.Add(ColorMapperFactory(
	//	func(col color.Color) bool { return true },
	//	func(col color.Color) color.Color {
	//		r, g, b, a := col.RGBA()
	//		return color.RGBA64{
	//			R: uint16(MaxUint16 - r),
	//			G: uint16(MaxUint16 - g),
	//			B: uint16(MaxUint16 - b),
	//			A: uint16(a),
	//		}
	//	}))
	cmr.Add(ColorMapperFactory(
		func(col color.Color) bool {
			r, g, b, _ := col.RGBA()
			return r+g+b < math.MaxUint16*3/2
		},
		func(col color.Color) color.Color {
			_, _, _, a := col.RGBA()
			return color.RGBA64{
				R: uint16(0),
				G: uint16(0),
				B: uint16(0),
				A: uint16(a),
			}
		}))

	cmr.Add(ColorMapperFactory(
		func(col color.Color) bool {
			r, g, b, _ := col.RGBA()
			return r+g+b > math.MaxUint16*3/2
		},
		func(col color.Color) color.Color {
			_, _, _, a := col.RGBA()
			return color.RGBA64{
				R: math.MaxUint16,
				G: math.MaxUint16,
				B: math.MaxUint16,
				A: uint16(a),
			}
		}))
	p := cmr.Apply(img)

	m := NewHorizontalMirror()
	m.TransferTo(ke)

	pm := m.Apply(p.Result())

	if err := Save(pm.Result(), "../resources/result1", "jpeg"); err != nil {
		fmt.Println(err.Error())
	}

	fishEye := NewFishEyeLens(Pt2D(350, 350), 200, 10)
	d := NewDistortion(filter.Extend, fishEye)
	d.TransferTo(ke)
	pd := d.Apply(img)

	if err := Save(pd.Result(), "../resources/result2", "jpeg"); err != nil {
		fmt.Println(err.Error())
	}

	dif := NewDifference(BinaryDifferenceFactory(0.01, color.Black, color.White))
	dif.TransferTo(ke)

	if err := Save(dif.Apply(img, pd.Result()).Result(), "../resources/dif", "jpeg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
