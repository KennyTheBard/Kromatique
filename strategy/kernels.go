package strategy

import (
	"math"

	"github.com/kennythebard/kromatique/utils"
)

func BoxBlurKernel(radius uint) utils.Kernel {
	size := int(radius*2 + 1)
	element := 1.0 / float64(size*size)
	kernel := make([][]float64, size)
	for i := 0; i < size; i++ {
		kernel[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			kernel[i][j] = element
		}
	}

	return kernel
}

func GaussianBlurKernel(sigma float64) utils.Kernel {
	if sigma <= 0 {
		return [][]float64{{1.0}}
	}

	radius := int(math.Ceil(sigma * 3.0))
	size := 2*radius + 1
	kernel := make([][]float64, 2*radius+1)

	calcElem := func(x, y, sigma float64) float64 {
		return math.Exp(-((x*x)+(y*y))/(2*sigma*sigma)) / (2 * sigma * sigma * math.Pi)
	}

	for i := 0; i < size; i++ {
		kernel[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			kernel[i][j] = calcElem(float64(j), float64(i), sigma)
		}
	}

	return kernel
}

func SobelKernels() []utils.Kernel {
	return []utils.Kernel{
		[][]float64{
			{1, 2, 1}, {0, 0, 0}, {-1, -2, -1},
		},
		[][]float64{
			{1, 0, -1}, {2, 0, -2}, {1, 0, -1},
		},
		//[][]float64{
		//	{-1, -2, -1}, {0, 0, 0}, {1, 2, 1},
		//},
		//[][]float64{
		//	{-1, 0, 1}, {-2, 0, 2}, {-1, 0, 1},
		//},
	}
}
