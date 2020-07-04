package core

import (
	"image"
	"sync"
)

type Promise func() image.Image

func PromiseImage(img image.Image, wg *sync.WaitGroup) Promise {
	return func() image.Image {
		wg.Wait()
		return img
	}
}
