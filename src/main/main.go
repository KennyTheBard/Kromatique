package main

import (
	. "./lib"
	. "./lib/geometry"
	. "./lib/imageio"
	. "./lib/morphing"
	. "./lib/render"
	. "./lib/strategy"
	"fmt"
	"image/color"
	"math"
)

func main() {
	img, err := Load("../resources/boat.png")
	if err != nil {
		panic(err)
	}

	ke := Parallel(4, 1000)
	defer ke.Stop()

	p := ke.Effect().ColorMapper([]MappingRule{
		func(in color.Color) color.Color {
			r, g, b, a := in.RGBA()
			avg := (r + g + b) / 3
			dev := uint32(math.Min(math.MaxUint16/4, math.Abs(float64(r)-float64(avg))))
			return color.RGBA64{
				R: uint16(r - dev),
				G: uint16(g),
				B: uint16(b),
				A: uint16(a),
			}
		},
	}).Apply(img)

	polygon := NewPolygon([]Point2D{
		Pt2D(10, 0),
		Pt2D(10, 10),
		Pt2D(20, 0),
		Pt2D(15, 15),
		Pt2D(0, 10),
	})
	polygon.Scale(10, 10).Rotate(math.Pi / 6)
	renderedImage := PolygonRender(polygon, MattePainter(color.RGBA{R: 255, A: 255}))

	if err := Save(renderedImage, "../resources/render", "png"); err != nil {
		fmt.Println(err.Error())
	}

	p = ke.Effect().Jitter(5).Apply(p.Result())
	if err := Save(p.Result(), "../resources/result", "png"); err != nil {
		fmt.Println(err.Error())
	}

	tri := BowyerWatson([]Point2D{
		Pt2D(10, 50),
		Pt2D(90, 50),
		Pt2D(50, 10),
		Pt2D(50, 90),
	}, Pt2D(0, 0), Pt2D(100, 100))
	for _, t := range tri.Triangles {
		fmt.Println(*t)
	}

}
