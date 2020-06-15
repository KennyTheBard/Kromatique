package main

import (
	. "./lib"
	. "./lib/imageio"
	. "./lib/morphing"
	. "./lib/strategy"
	"fmt"
)

func main() {
	img, err := Load("../resources/boat.png")
	if err != nil {
		panic(err)
	}

	ke := Parallel(4, 1000)
	defer ke.Stop()

	img = ke.Effect().ColorMapper([]MappingRule{
		BrightnessMapping(0.25),
	}).Apply(img).Result()

	mesh := NewMesh(img)
	BowyerWatson(mesh, []Vertex{
		Vx(100, 100),
		Vx(250, 250),
	})
	morph := NewMeshDeformation(mesh, img, map[Vertex]Vertex{
		Vx(100, 100): Vx(250, 250),
		Vx(250, 250): Vx(350, 200),
	}, 1)

	if err := Save(morph.Deform(), "../resources/result", "png"); err != nil {
		fmt.Println(err.Error())
	}

}
