package lib

func Lerp(A, B, alpha float64) float64 {
	return (1 - alpha) * A + alpha * B
}
