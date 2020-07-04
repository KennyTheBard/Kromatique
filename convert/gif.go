package convert

import (
	"errors"
	"image"
	"image/draw"
	"image/gif"

	"github.com/kennythebard/kromatique/core"
	"github.com/kennythebard/kromatique/utils"
)

func ImageToGif(images []image.Image, delays []int, colorPalletSize int) (*gif.GIF, error) {
	if len(images) != len(delays) {
		return nil, errors.New("images and delays have different lengths")
	}

	if colorPalletSize < 1 {
		return nil, errors.New("A pallet cannot be without any color")
	}

	palletedImages := make([]*image.Paletted, len(images))
	core.Parallelize(len(images), func(curr int) {
		palletedImages[curr] = image.NewPaletted(images[curr].Bounds(), utils.GeneratePallet(images[curr], colorPalletSize))
		draw.Draw(palletedImages[curr], palletedImages[curr].Rect, images[curr], images[curr].Bounds().Min, draw.Over)
	})

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
