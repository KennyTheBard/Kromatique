package lib

import (
	"./blend"
	"./convert"
	"./core"
	"./effect"
	"./histogram"
	"./morphing"
	"image"
)

type Krom struct {
	engine    core.Engine
	blend     *blend.Factory
	convert   *convert.Factory
	effect    *effect.Factory
	histogram *histogram.Factory
}

func (krom *Krom) Blend() *blend.Factory {
	return krom.blend
}

func (krom *Krom) Morph(src, dst image.Image, srcPoints, dstPoints []morphing.Vertex, numSteps int) []image.Image {
	return morphing.Morph(krom.engine, src, dst, srcPoints, dstPoints, numSteps)
}

func (krom *Krom) Convert() *convert.Factory {
	return krom.convert
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

func Parallel(numWorkers, queueSize int) *Krom {
	krom := new(Krom)
	krom.engine = core.NewPoolEngine(numWorkers, queueSize)
	krom.blend = blend.NewFactory(krom.engine)
	krom.convert = convert.NewFactory(krom.engine)
	krom.effect = effect.NewFactory(krom.engine)
	krom.histogram = histogram.NewFactory(krom.engine)

	return krom
}

func Sequential() *Krom {
	krom := new(Krom)
	krom.engine = core.SequentialEngine{}
	krom.effect = effect.NewFactory(krom.engine)
	krom.convert = convert.NewFactory(krom.engine)
	krom.blend = blend.NewFactory(krom.engine)
	krom.histogram = histogram.NewFactory(krom.engine)

	return krom
}
