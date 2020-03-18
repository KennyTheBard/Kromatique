package utils

import "math"

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}

	if x > max {
		return max
	}

	return x
}

// ClampUint8 returns given value clamped into uint8 range
func ClampUint8(x float64) float64 {
	return Clamp(x, 0, math.MaxUint8)
}

// ClampUint16 returns given value clamped into uint16 range
func ClampUint16(x float64) float64 {
	return Clamp(x, 0, math.MaxUint16)
}

// ClampUint32 returns given value clamped into uint32 range
func ClampUint32(x float64) float64 {
	return Clamp(x, 0, math.MaxUint32)
}

// ClampUint64 returns given value clamped into uint64 range
func ClampUint64(x float64) float64 {
	return Clamp(x, 0, math.MaxUint64)
}
