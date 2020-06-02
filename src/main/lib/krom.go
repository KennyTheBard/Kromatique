package lib

import (
	"./blend"
	"./core"
	"./effect"
)

type Krom struct {
	engine core.Engine
	effect *effect.Factory
	blend  *blend.Factory
}

func (krom *Krom) Effect() *effect.Factory {
	return krom.effect
}

func (krom *Krom) Blend() *blend.Factory {
	return krom.blend
}

func (krom *Krom) Stop() {
	krom.engine.Stop()
}

func Parallel(workForce, queueSize int) *Krom {
	krom := new(Krom)
	krom.engine = core.NewPoolEngine(workForce, queueSize)
	krom.effect = effect.NewFactory(krom.engine)
	krom.blend = blend.NewFactory(krom.engine)

	return krom
}

func Sequential(workForce, queueSize int) *Krom {
	krom := new(Krom)
	krom.engine = core.SequentialEngine{}
	krom.effect = effect.NewFactory(krom.engine)
	krom.blend = blend.NewFactory(krom.engine)

	return krom
}
