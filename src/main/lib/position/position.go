package position

import "image"

type Position interface {
	Find(image.Rectangle) image.Point
}

type MixedPosition struct {
	posX, posY AxialPosition
}

func (pos *MixedPosition) Find(bounds image.Rectangle) image.Point {
	return image.Pt(pos.posX.Find(bounds.Min.X, bounds.Max.X), pos.posY.Find(bounds.Min.Y, bounds.Max.Y))
}

func Pos(posX, posY AxialPosition) *MixedPosition {
	pos := new(MixedPosition)
	pos.posX = posX
	pos.posY = posY

	return pos
}
