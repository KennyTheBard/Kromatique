package main

import (
	lib "./lib"
	"fmt"
)

func main() {
	img := lib.Load("../resources/test.jpg")

	ke := lib.NewKromEngine(10, 0)
	p := ke.Sepia().Apply(img)

	if err := lib.Save(p.Result(), "../resources/result", "jpg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
