package lib

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
func Load(path string) image.Image {
	file, errOpen := os.Open(path)
	if errOpen != nil {
		return nil
	}
	defer file.Close()

	img, _, errDecode := image.Decode(file)
	if errDecode != nil {
		return nil
	}
	return img
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

// Create returns an image with the given size
func Create(dimension image.Rectangle) image.Image {
	img := image.NewRGBA(dimension)
	return img
}

// CreateBackground returns an image with the given size and a monochrome background of the given color
func CreateBackground(dimension image.Rectangle, col color.Color) image.Image {
	img := image.NewRGBA(dimension)
	draw.Draw(img, img.Bounds(), &image.Uniform{C: col}, dimension.Min, draw.Src)
	return img
}
