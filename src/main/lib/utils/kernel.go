package utils

type Kernel [][]float64

func (m *Kernel) Get(x, y int) float64 {
	return (*m)[y][x]
}

func (m *Kernel) Set(x, y int, val float64) {
	(*m)[y][x] = val
}

func (m *Kernel) Radius() int {
	return len(*m) / 2
}
