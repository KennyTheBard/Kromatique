package utils

// Kernel encapsulates the data for a matrix
type Kernel [][]float64

// Get returns the value on the line y, row x
func (m *Kernel) Get(x, y int) float64 {
	return (*m)[y][x]
}

// Radius returns the radius of the Kernel
func (m *Kernel) Radius() int {
	return len(*m) / 2
}
