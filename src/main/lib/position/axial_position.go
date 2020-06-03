package position

import "math"

type AxialPosition interface {
	Find(min, max int) int
}

type Anchor int

const (
	Start Anchor = 0
	End   Anchor = 1 << iota
)

type FixedPosition struct {
	anchor Anchor
	value  int
}

func (pos *FixedPosition) Find(min, max int) int {
	if pos.anchor == Start {
		return min + pos.value
	} else {
		return max - pos.value
	}
}

func (pos *FixedPosition) Anchor(anchor Anchor) *FixedPosition {
	pos.anchor = anchor
	return pos
}

func (pos *FixedPosition) Value(value int) *FixedPosition {
	pos.value = value
	return pos
}

func Fixed(value int) *FixedPosition {
	pos := new(FixedPosition)
	pos.value = value

	return pos
}

type PercentPosition struct {
	anchor Anchor
	value  float64
}

func (pos *PercentPosition) Find(min, max int) int {
	d := max - min
	if pos.anchor == Start {
		return int(math.Min(float64(min)+math.Round(float64(d)*pos.value), float64(max)))
	} else {
		return int(math.Max(float64(max)-math.Round(float64(d)*pos.value), float64(min)))
	}
}

func (pos *PercentPosition) Anchor(anchor Anchor) *PercentPosition {
	pos.anchor = anchor
	return pos
}

func (pos *PercentPosition) Value(value float64) *PercentPosition {
	pos.value = value
	return pos
}

func Percent(value float64) *PercentPosition {
	pos := new(PercentPosition)
	pos.value = value

	return pos
}
