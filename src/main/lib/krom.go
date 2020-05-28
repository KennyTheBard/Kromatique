package lib

import (
	"./core"
	"./effect"
)

type Krom struct {
	engine  core.Engine
	factory *effect.Factory
}

func (krom *Krom) Effect() *effect.Factory {
	return krom.factory
}

func (krom *Krom) Stop() {
	krom.engine.Stop()
}

func Parallel(workForce, queueSize int) *Krom {
	krom := new(Krom)
	krom.engine = core.NewPoolEngine(workForce, queueSize)
	krom.factory = effect.NewFactory(krom.engine)

	return krom
}

func Sequential(workForce, queueSize int) *Krom {
	krom := new(Krom)
	krom.engine = core.SequentialEngine{}
	krom.factory = effect.NewFactory(krom.engine)

	return krom
}
