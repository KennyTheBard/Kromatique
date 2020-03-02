package utils

const MaxUint8 = 2<<7 - 1
const MaxUint16 = 2<<15 - 1
const MaxUint32 = 2<<31 - 1
const MaxUint64 = 2<<63 - 1

const MinUint8 = 0
const MinUint16 = 0
const MinUint32 = 0
const MinUint64 = 0

// ClampUint8 returns given value clamped into uint8 range
func ClampUint8(x float64) float64 {
	if x < 0 {
		return MinUint8
	}

	if x > MaxUint8 {
		return MaxUint8
	}

	return x
}

// ClampUint16 returns given value clamped into uint16 range
func ClampUint16(x float64) float64 {
	if x < 0 {
		return MinUint16
	}

	if x > MaxUint16 {
		return MaxUint16
	}

	return x
}

// ClampUint32 returns given value clamped into uint32 range
func ClampUint32(x float64) float64 {
	if x < 0 {
		return MinUint32
	}

	if x > MaxUint32 {
		return MaxUint32
	}

	return x
}

// ClampUint64 returns given value clamped into uint64 range
func ClampUint64(x float64) float64 {
	if x < 0 {
		return MinUint64
	}

	if x > MaxUint64 {
		return MaxUint64
	}

	return x
}
