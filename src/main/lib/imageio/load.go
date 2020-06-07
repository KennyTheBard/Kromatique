package imageio

import (
	"image"
	"os"
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
