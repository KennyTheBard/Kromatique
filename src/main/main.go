package main

import (
	"fmt"
	"image/color"

	. "./lib"
	. "./lib/effect/distortion"
	. "./lib/effect/filter"
	. "./lib/effect/mapper"
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
			return r+g+b < MaxUint16*3/2
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
			return r+g+b > MaxUint16*3/2
		},
		func(col color.Color) color.Color {
			_, _, _, a := col.RGBA()
			return color.RGBA64{
				R: MaxUint16,
				G: MaxUint16,
				B: MaxUint16,
				A: uint16(a),
			}
		}))

	p := cmr.Apply(img)
	if err := Save(p.Result(), "../resources/result1", "jpeg"); err != nil {
		fmt.Println(err.Error())
	}

	test := NewFishEyeLens(Pt2D(3, 3), 2, 0.5)
	vm := test.VectorMap()
	fmt.Printf("(%d, %d) to (%d, %d)\n", vm.Bounds().Min.X, vm.Bounds().Min.Y, vm.Bounds().Max.X, vm.Bounds().Max.Y)
	for y := vm.Bounds().Min.Y; y < vm.Bounds().Max.Y; y++ {
		for x := vm.Bounds().Min.X; x < vm.Bounds().Max.X; x++ {
			fmt.Printf("(%.2f, %.2f) ", vm.At(x, y).X, vm.At(x, y).Y)
		}
		fmt.Println()
	}

	fishEye := NewFishEyeLens(Pt2D(350, 350), 200, 10)
	d := NewDistortion(Extend, fishEye)
	d.TransferTo(ke)

	if err := Save(d.Apply(img).Result(), "../resources/result2", "jpeg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
