package position

import "math"

// AxialPosition is in interface used to map an abstract value
// to a concrete value for an axial position
type AxialPosition interface {
	// Find returns the mapped encapsulated position on the given interval
	Find(min, max int) int
}

// FixedPosition is an implementation for AxialPosition
// that calculates a fixed position with an offset
type FixedPosition struct {
	reversed bool
	offset   int
}

// Find returns the mapped position on the given interval based on offset
// if reversed flag is set to false or negative offset otherwise
func (pos *FixedPosition) Find(min, max int) int {
	if pos.reversed {
		return min - pos.offset
	} else {
		return max + pos.offset
	}
}

// Reverse sets the reversed flag to the given value
func (pos *FixedPosition) Reverse(reversed bool) *FixedPosition {
	pos.reversed = reversed
	return pos
}

// Offset sets the offset to the given value
func (pos *FixedPosition) Offset(offset int) *FixedPosition {
	pos.offset = offset
	return pos
}

// Fixed returns a FixedPosition with the given offset
// and reversed flag set to false
func Fixed(offset int) *FixedPosition {
	pos := new(FixedPosition)
	pos.offset = offset

	return pos
}

// PercentPosition is an implementation for AxialPosition
// that calculates a fixed position with an offset
type PercentPosition struct {
	reversed bool
	percent  float64
}

// Find returns the mapped position on the given interval based on percent
// if reversed flag is false or negative percent otherwise
func (pos *PercentPosition) Find(min, max int) int {
	d := max - min
	if pos.reversed {
		return int(math.Max(float64(max)-math.Round(float64(d)*pos.percent), float64(min)))
	} else {
		return int(math.Min(float64(min)+math.Round(float64(d)*pos.percent), float64(max)))
	}
}

// Reverse sets the reversed flag to the given value
func (pos *PercentPosition) Reverse(reversed bool) *PercentPosition {
	pos.reversed = reversed
	return pos
}

// Percent sets the percent to the given value
func (pos *PercentPosition) Percent(percent float64) *PercentPosition {
	pos.percent = percent
	return pos
}

// Percent returns a PercentPosition with the given percent
// and reversed flag set to false
func Percent(percent float64) *PercentPosition {
	pos := new(PercentPosition)
	pos.percent = percent

	return pos
}
