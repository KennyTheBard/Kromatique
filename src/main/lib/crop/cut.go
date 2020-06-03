package crop

import (
	"image"
	"image/color"

	"../position"
)

type ImageSlice struct {
	origin    image.Image
	cutBounds image.Rectangle
}

func (slice *ImageSlice) ColorModel() color.Model {
	return slice.origin.ColorModel()
}

func (slice *ImageSlice) Bounds() image.Rectangle {
	return slice.cutBounds
}

func (slice *ImageSlice) At(x, y int) color.Color {
	if image.Pt(x, y).In(slice.cutBounds) {
		return slice.origin.At(x, y)
	}

	return color.RGBA{}
}

func Cut(img image.Image, start, end position.Position) image.Image {
	slice := new(ImageSlice)
	slice.origin = img

	boundsStart := start.Find(img.Bounds())
	boundsEnd := end.Find(img.Bounds())
	slice.cutBounds = image.Rect(boundsStart.X, boundsStart.Y, boundsEnd.X, boundsEnd.Y)

	return slice
}
