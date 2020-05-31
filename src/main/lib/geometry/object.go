package geometry

type IObject2D interface {
	Translate(float64, float64)
	Scale(float64, float64)
	Rotate(float64)
}

type Object2D struct {
	model, inverse ModelMatrix2D
}

func (object *Object2D) Translate(x, y float64) {
	object.model.Translate(x, y)
	object.inverse = object.inverse.Multiply(TranslateMatrix(-x, -y))
}

func (object *Object2D) Scale(x, y float64) {
	object.model.Scale(x, y)
	object.inverse = object.inverse.Multiply(ScaleMatrix(1/x, 1/y))
}

func (object *Object2D) Rotate(a float64) {
	object.model.Rotate(a)
	object.inverse = object.inverse.Multiply(RotateMatrix(-a))
}
