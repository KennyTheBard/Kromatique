package effect

import (
	"fmt"
	"image"
	"image/draw"
	"math/rand"
	"time"

	"../core"
	"../utils"
)

// Jitter serves as a generic customizable structure that encapsulates
// the logic needed to apply a jitter effect on an image
type Jitter struct {
	engine core.Engine
	radius int
}

func (effect *Jitter) Apply(img image.Image) *core.Promise {
	ret := utils.CreateRGBA(img.Bounds())
	contract := effect.engine.Contract(img.Bounds().Dy())
	rand.Seed(time.Now().Unix())
	randCoordinate := func() int {
		return rand.Intn(effect.radius*2) - effect.radius
	}

	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				offsetX, offsetY := randCoordinate(), randCoordinate()
				newX := utils.Max(utils.Min(x+offsetX, img.Bounds().Max.X), img.Bounds().Min.X)
				newY := utils.Max(utils.Min(y+offsetY, img.Bounds().Max.Y), img.Bounds().Min.Y)

				ret.(draw.Image).Set(x, y, img.At(newX, newY))
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}
