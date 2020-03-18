package distorsion

import (
	"image"
)

type LensOperation func(Vector, Vector) Vector

func Add(a, b Vector) Vector {
	return Vec(a.X+b.X, a.Y+b.Y)
}

func Diff(a, b Vector) Vector {
	return Vec(a.X-b.X, a.Y-b.Y)
}

func Merge(a, b Vector) Vector {
	return Vec((a.X+b.X)/2, (a.Y+b.Y)/2)
}

type LensComposer struct {
	op          LensOperation
	left, right Lens
}

func (c LensComposer) VectorMap() VectorMap {
	a := c.left.VectorMap()
	b := c.right.VectorMap()
	bounds := a.bounds.Union(b.bounds)
	vm := NewVectorMap(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			p := image.Pt(x, y)
			var va, vb Vector
			if p.In(a.Bounds()) {
				va = a.At(x, y)
			}
			if p.In(b.Bounds()) {
				vb = b.At(x, y)
			}

			vm.Set(x, y, c.op(va, vb))
		}
	}

	return vm
}
