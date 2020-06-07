package imageio

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

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
