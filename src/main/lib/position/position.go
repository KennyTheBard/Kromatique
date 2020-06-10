package position

import "image"

// Position is an interface for an abstract point
// that can be mapped to a image.Point of a image.Rectangle
type Position interface {
	Find(image.Rectangle) image.Point
}

// MixedPosition is an implementation of the Position interface
// that can encapsulates a different method of mapping for each axis
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
