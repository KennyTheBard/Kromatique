package geometry

type IObject2D interface {
	Translate(float64, float64) *Object2D
	Scale(float64, float64) *Object2D
	Rotate(float64) *Object2D
}

type Object2D struct {
	model, inverse ModelMatrix2D
}

func (object *Object2D) Model() *ModelMatrix2D {
	return &object.model
}

func (object *Object2D) Inverse() *ModelMatrix2D {
	return &object.inverse
}

func (object *Object2D) Translate(x, y float64) *Object2D {
	object.Model().Translate(x, y)
	object.inverse = object.Inverse().Multiply(TranslateMatrix(-x, -y))
	return object
}

func (object *Object2D) Scale(x, y float64) *Object2D {
	object.Model().Scale(x, y)
	object.inverse = object.Inverse().Multiply(ScaleMatrix(1/x, 1/y))
	return object
}

func (object *Object2D) Rotate(a float64) *Object2D {
	object.Model().Rotate(-a)
	object.inverse = object.Inverse().Multiply(RotateMatrix(a))
	return object
}

func InitObject(object *Object2D) {
	object.model = Eye2D()
	object.inverse = Eye2D()
}
