package position

import "math"

// AxialPosition is in interface used to map an abstract value
// to a concrete value for an axial position
type AxialPosition interface {
	Find(min, max int) int
}

type FixedPosition struct {
	reverse bool
	value   int
}

func (pos *FixedPosition) Find(min, max int) int {
	if pos.reverse {
		return min - pos.value
	} else {
		return max + pos.value
	}
}

func (pos *FixedPosition) Reverse(reverse bool) *FixedPosition {
	pos.reverse = reverse
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
	reverse bool
	value   float64
}

func (pos *PercentPosition) Find(min, max int) int {
	d := max - min
	if pos.reverse {
		return int(math.Max(float64(max)-math.Round(float64(d)*pos.value), float64(min)))
	} else {
		return int(math.Min(float64(min)+math.Round(float64(d)*pos.value), float64(max)))
	}
}

func (pos *PercentPosition) Reverse(reverse bool) *PercentPosition {
	pos.reverse = reverse
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
