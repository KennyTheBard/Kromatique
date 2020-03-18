package distorsion

import "image"

type Vector struct {
	X, Y float64
}

func Vec(x, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

func VecZero() Vector {
	return Vec(0, 0)
}

type VectorMap struct {
	vs     [][]Vector
	bounds image.Rectangle
}

func (m VectorMap) At(x, y int) Vector {
	return m.vs[y-m.bounds.Min.Y][x-m.bounds.Min.X]
}

func (m *VectorMap) Set(x, y int, v Vector) {
	m.vs[y-m.bounds.Min.Y][x-m.bounds.Min.X] = v
}

func (m VectorMap) Bounds() image.Rectangle {
	return m.bounds
}

func NewVectorMap(bounds image.Rectangle) VectorMap {
	vs := make([][]Vector, bounds.Dy())
	for i := 0; i < bounds.Dy(); i++ {
		vs[i] = make([]Vector, bounds.Dx())
	}

	return VectorMap{vs: vs, bounds: bounds}
}
