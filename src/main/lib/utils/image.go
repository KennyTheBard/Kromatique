package utils

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// Load returns an image object of the file located at given path
func Load(path string) (image.Image, error) {
	file, errOpen := os.Open(path)
	if errOpen != nil {
		return nil, errOpen
	}
	defer file.Close()

	img, _, errDecode := image.Decode(file)
	if errDecode != nil {
		return nil, errDecode
	}
	return img, nil
}

// Save write the given image in a file of chosen format at the given path
func Save(img image.Image, path, format string) error {
	file, errCreate := os.Create(strings.Join([]string{path, format}, "."))
	if errCreate != nil {
		return errCreate
	}
	defer file.Close()

	switch format {
	case "png":
		return png.Encode(file, img)

	case "jpeg", "jpg":
		return jpeg.Encode(file, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}

	return errors.New("Unsupported file format: " + format)
}

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
