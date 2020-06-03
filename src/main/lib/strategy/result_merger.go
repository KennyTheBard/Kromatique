package strategy

import "math"

// ResultMerger defines an interface for all functions used
// to determine the merged result of multiple applied filters for MultiKernel
type ResultMerger func([]float64) float64

// SobelGradient defines the rule used to merge results for sobel operator;
// despite sobel using only 2 kernels, the function handles more values
func SobelGradient(results []float64) float64 {
	var sum float64
	for _, res := range results {
		sum += res * res
	}

	return math.Sqrt(sum)
}
