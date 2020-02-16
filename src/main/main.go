package main

import (
	lib "./lib"
	"fmt"
)

func main() {
	img := lib.Load("../resources/test.jpg")

	ke := lib.NewKromEngine(10, 0)
	ret := ke.Grayscale().Apply(img)

	if err := lib.Save(ret, "../resources/result", "jpg"); err != nil {
		fmt.Println(err.Error())
	}

	ke.Stop()
}
