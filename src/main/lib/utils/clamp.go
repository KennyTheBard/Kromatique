package utils

const MaxUint8 = 2<<7 - 1
const MaxUint16 = 2<<15 - 1
const MaxUint32 = 2<<31 - 1
const MaxUint64 = 2<<63 - 1

const MinUint8 = 0
const MinUint16 = 0
const MinUint32 = 0
const MinUint64 = 0

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
	return Clamp(x, MinUint8, MaxUint8)
}

// ClampUint16 returns given value clamped into uint16 range
func ClampUint16(x float64) float64 {
	return Clamp(x, MinUint16, MaxUint16)
}

// ClampUint32 returns given value clamped into uint32 range
func ClampUint32(x float64) float64 {
	return Clamp(x, MinUint32, MaxUint32)
}

// ClampUint64 returns given value clamped into uint64 range
func ClampUint64(x float64) float64 {
	return Clamp(x, MinUint64, MaxUint64)
}
