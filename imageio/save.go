package imageio

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// Save write the given image in a file of chosen format at the given path
func Save(img interface{}, path, format string) error {
	file, errCreate := os.Create(strings.Join([]string{path, format}, "."))
	if errCreate != nil {
		return errCreate
	}
	defer file.Close()

	switch strings.ToLower(format) {
	case "png":
		return png.Encode(file, img.(image.Image))

	case "jpeg", "jpg":
		return jpeg.Encode(file, img.(image.Image), &jpeg.Options{Quality: jpeg.DefaultQuality})

	case "gif":
		return gif.EncodeAll(file, img.(*gif.GIF))

	default:
		return errors.New("Unsupported file format: " + format)
	}
}
