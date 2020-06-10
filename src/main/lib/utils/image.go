package utils

import (
	"image"
	"image/color"
	"image/draw"
)

// CreateRGBA returns an image with the given size
func CreateRGBA(bounds image.Rectangle) image.Image {
	img := image.NewRGBA(bounds)
	return img
}

// CreateBackground returns an image with the given size and a monochrome background of the given color
func CreateBackground(bounds image.Rectangle, col color.Color) image.Image {
	img := image.NewRGBA(bounds)
	draw.Draw(img, img.Bounds(), &image.Uniform{C: col}, bounds.Min, draw.Src)
	return img
}
