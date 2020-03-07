package filter

import "math"

type ResultMergingStrategy func([]float64) float64

func SobelGradient(results []float64) float64 {
	var sum float64
	for _, res := range results {
		sum += res * res
	}

	return math.Sqrt(sum)
}
