package utils

type Matrix [][]float64

func (m *Matrix) Get(x, y int) float64 {
	return (*m)[y][x]
}

func (m *Matrix) Set(x, y int, val float64) {
	(*m)[y][x] = val
}

func (m *Matrix) Radius() int {
	return (len(*m) - 1) / 2
}
