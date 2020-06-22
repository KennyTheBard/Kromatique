package convert

import (
	"errors"
	"image"
	"image/draw"
	"image/gif"

	"../core"
	"../utils"
)

func ImageToGif(engine core.Engine, images []image.Image, delays []int, colorPalletSize int) (*gif.GIF, error) {
	if len(images) != len(delays) {
		return nil, errors.New("images and delays have different lengths")
	}

	if colorPalletSize < 1 {
		return nil, errors.New("A pallet cannot be without any color")
	}

	contract := engine.Contract()
	palletedImages := make([]*image.Paletted, len(images))
	for idx := range images {
		curr := idx
		contract.PlaceOrder(func() {
			palletedImages[curr] = image.NewPaletted(images[curr].Bounds(), utils.GeneratePallet(images[curr], colorPalletSize))
			draw.Draw(palletedImages[curr], palletedImages[curr].Rect, images[curr], images[curr].Bounds().Min, draw.Over)
		})
	}

	contract.Deadline()

	gifImage := &gif.GIF{
		Image:           palletedImages,
		Delay:           delays,
		LoopCount:       0,
		Disposal:        nil,
		Config:          image.Config{},
		BackgroundIndex: 0,
	}

	return gifImage, nil
}
