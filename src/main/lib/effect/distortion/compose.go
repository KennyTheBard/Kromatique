package distorsion

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

func (c LensComposer) VecAt(x, y int) Vector {
	return c.op(c.left.VecAt(x, y), c.right.VecAt(x, y))
}
