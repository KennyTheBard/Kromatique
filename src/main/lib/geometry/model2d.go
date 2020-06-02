package geometry

import "math"

type ModelMatrix2D [3][3]float64

func (model *ModelMatrix2D) Apply(p Point2D) Point2D {
	return Pt2D(p.X*model[0][0]+p.Y*model[0][1]+model[0][2], p.X*model[1][0]+p.Y*model[1][1]+model[1][2])
}

func (model *ModelMatrix2D) Multiply(other *ModelMatrix2D) ModelMatrix2D {
	ret := Eye2D()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ret[i][j] = 0
			for k := 0; k < 3; k++ {
				ret[i][j] += model[i][k] * other[k][j]
			}
		}
	}

	return ret
}

func (model *ModelMatrix2D) Translate(x, y float64) {
	*model = TranslateMatrix(x, y).Multiply(model)
}

func (model *ModelMatrix2D) Scale(x, y float64) {
	*model = ScaleMatrix(x, y).Multiply(model)
}

func (model *ModelMatrix2D) Rotate(a float64) {
	*model = RotateMatrix(a).Multiply(model)
}

func Eye2D() ModelMatrix2D {
	return ModelMatrix2D{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
}

func TranslateMatrix(tx, ty float64) *ModelMatrix2D {
	return &ModelMatrix2D{{1, 0, tx}, {0, 1, ty}, {0, 0, 1}}
}

func ScaleMatrix(sx, sy float64) *ModelMatrix2D {
	return &ModelMatrix2D{{sx, 0, 0}, {0, sy, 0}, {0, 0, 1}}
}

func RotateMatrix(a float64) *ModelMatrix2D {
	return &ModelMatrix2D{{math.Cos(a), -math.Sin(a), 0}, {math.Sin(a), math.Cos(a), 0}, {0, 0, 1}}
}

func MM2D(p Point2D) ModelMatrix2D {
	return ModelMatrix2D{{0, 0, p.X}, {0, 0, p.Y}, {0, 0, 1}}
}
