package convert

import (
	"../core"
	"image"
	"image/gif"
)

type Factory struct {
	engine core.Engine
}

func NewFactory(engine core.Engine) *Factory {
	f := new(Factory)
	f.engine = engine

	return f
}

func (f Factory) ToGif(images []image.Image, delays []int, colorPalletSize int) (*gif.GIF, error) {
	return ImageToGif(f.engine, images, delays, colorPalletSize)
}
