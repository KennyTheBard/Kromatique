package main

import (
	"fmt"
	"image/color"

	. "./lib"
	. "./lib/effect/mapper"
	. "./lib/utils"
)

func main() {
	img := Load("../resources/test.jpg")

	ke := NewKromEngine(10, 0)

	cmr := NewColorMapperRunner()
	cmr.TransferTo(ke)
	cmr.Add(ColorMapperFactory(
		func(col color.Color) bool { return true },
		func(col color.Color) color.Color {
			r, g, b, a := col.RGBA()
			return color.RGBA64{
				R: uint16(MaxUint16 - r),
				G: uint16(MaxUint16 - g),
				B: uint16(MaxUint16 - b),
				A: uint16(a),
			}
		}))

	cmr.Add(ColorMapperFactory(
		func(col color.Color) bool {
			r, g, _, _ := col.RGBA()
			return r < g
		},
		func(col color.Color) color.Color {
			_, g, b, a := col.RGBA()
			return color.RGBA64{
				R: uint16(g),
				G: uint16(g),
				B: uint16(b),
				A: uint16(a),
			}
		}))

	cmr.Add(ColorMapperFactory(
		func(col color.Color) bool {
			r, _, b, _ := col.RGBA()
			return r < b
		},
		func(col color.Color) color.Color {
			_, g, b, a := col.RGBA()
			return color.RGBA64{
				R: uint16(b),
				G: uint16(g),
				B: uint16(b),
				A: uint16(a),
			}
		}))

	p := cmr.Apply(img)

	if err := Save(p.Result(), "../resources/result", "jpeg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
