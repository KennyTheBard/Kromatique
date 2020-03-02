package utils

// LERP is a simple linear interpolation function
func LERP(A, B, alpha float64) float64 {
	return (1-alpha)*A + alpha*B
}
