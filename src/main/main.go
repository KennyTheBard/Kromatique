package main

import (
	"fmt"
	"image/color"
	"math"

	. "./lib"
	. "./lib/effect/flip"
	. "./lib/effect/mapper"
	. "./lib/utils"
)

func main() {
	img := Load("../resources/test.jpg")

	ke := NewKromEngine(10, 0)

	cmr := NewColorMapperRunner()
	cmr.TransferTo(ke)
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

	if err := Save(pm.Result(), "../resources/result", "jpeg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
