package lib

import (
	"./blend"
	"./core"
	"./effect"
	"./histogram"
)

type Krom struct {
	engine    core.Engine
	blend     *blend.Factory
	effect    *effect.Factory
	histogram *histogram.Factory
}

func (krom *Krom) Blend() *blend.Factory {
	return krom.blend
}

func (krom *Krom) Effect() *effect.Factory {
	return krom.effect
}

func (krom *Krom) Histo() *histogram.Factory {
	return krom.histogram
}

func (krom *Krom) Stop() {
	krom.engine.Stop()
}

func Parallel(workForce, queueSize int) *Krom {
	krom := new(Krom)
	krom.engine = core.NewPoolEngine(workForce, queueSize)
	krom.blend = blend.NewFactory(krom.engine)
	krom.effect = effect.NewFactory(krom.engine)
	krom.histogram = histogram.NewFactory(krom.engine)

	return krom
}

func Sequential(workForce, queueSize int) *Krom {
	krom := new(Krom)
	krom.engine = core.SequentialEngine{}
	krom.effect = effect.NewFactory(krom.engine)
	krom.blend = blend.NewFactory(krom.engine)

	return krom
}
