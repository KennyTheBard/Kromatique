package crop

import (
	"image"
	"image/draw"

	"../position"
	"../utils"
)

func Copy(img image.Image, start, end position.Position) image.Image {
	boundsStart := start.Find(img.Bounds())
	boundsEnd := end.Find(img.Bounds())

	bounds := image.Rect(boundsStart.X, boundsStart.Y, boundsEnd.X, boundsEnd.Y)
	imgCopy := utils.CreateRGBA(bounds)

	draw.Draw(imgCopy.(draw.Image), bounds, img, boundsStart, draw.Over)

	return imgCopy
}
