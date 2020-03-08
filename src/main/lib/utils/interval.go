package utils

// ColorInterval encapsulates a closed discontinuous interval
// of integer numbers representing color values
type ColorInterval struct {
	min, max uint32
}

func (i ColorInterval) Min() uint32 {
	return i.min
}

func (i ColorInterval) Max() uint32 {
	return i.max
}

func (i ColorInterval) Contains(val uint32) bool {
	return i.min <= val && val <= i.max
}

func NewColorInterval(min, max uint32) *ColorInterval {
	i := new(ColorInterval)
	i.min = min
	i.max = max

	return i
}
