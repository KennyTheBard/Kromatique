package lib

import "image"

// ImageEffect is the core interface of all image processing effects
type ImageEffect interface {
	Apply(image.Image) Promise
}

type BaseEffect struct {
	engine *KromEngine
}

func (base *BaseEffect) TransferTo(engine *KromEngine) {
	base.engine = engine
}

func (base *BaseEffect) GetEngine() *KromEngine {
	return base.engine
}