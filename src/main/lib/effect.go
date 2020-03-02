package lib

import "image"

// ImageEffect is the core interface of all image processing effects
type ImageEffect interface {
	Apply(image.Image) Promise
}

// BaseEffect is a structure useful to ensure compatibility with the library
// as easy as possible, encapsulating interactions with the engine in use
type BaseEffect struct {
	engine *KromEngine
}

// TransferTo is used to change the engine currently in use
func (base *BaseEffect) TransferTo(engine *KromEngine) {
	base.engine = engine
}

// GetEngine simple getter to access the engine currently in use
func (base *BaseEffect) GetEngine() *KromEngine {
	return base.engine
}
