package position

import (
	"image"
)

// Position is an interface that serves for exposing generic ways to
// decide upon the position of a point in an image depending of resolution
type Position interface {
	Get(image.Rectangle) image.Point
}

// FixedPosition is a simple wrapper on image.Point
type FixedPosition struct {
	image.Point
}

func (p *FixedPosition) Get(_ image.Rectangle) image.Point {
	return p.Point
}

func NewFixedPosition(x, y int) *FixedPosition {
	fp := new(FixedPosition)
	fp.X = x
	fp.Y = y

	return fp
}

// RelativePosition is a position representation that can use two different
// representations for each axis in order to achieve greater flexibility
type RelativePosition struct {
	x, y *RelativeAxialPosition
}

func (p *RelativePosition) Get(bounds image.Rectangle) image.Point {
	return image.Pt(p.x.Get(bounds.Min.X, bounds.Max.X), p.y.Get(bounds.Min.Y, bounds.Max.Y))
}

func NewRelativePosition(x, y *RelativeAxialPosition) *RelativePosition {
	rp := new(RelativePosition)
	rp.x = x
	rp.y = y

	return rp
}
