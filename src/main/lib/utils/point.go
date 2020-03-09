package utils

type Point2D struct {
	X, Y float64
}

func NewPoint2D(X, Y float64) *Point2D {
	p := new(Point2D)
	p.X = X
	p.Y = Y

	return p
}
