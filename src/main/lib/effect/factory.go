package effect

import (
	"../core"
	"../strategy"
	"../utils"
)

type Factory struct {
	engine core.Engine
}

func NewFactory(engine core.Engine) *Factory {
	f := new(Factory)
	f.engine = engine

	return f
}

func (f Factory) HFlip() *Flip {
	effect := new(Flip)
	effect.engine = f.engine
	effect.strategy = HorizontalFlip

	return effect
}

func (f Factory) VFlip() *Flip {
	effect := new(Flip)
	effect.engine = f.engine
	effect.strategy = VerticalFlip

	return effect
}

func (f Factory) Filter(edgeHandling strategy.EdgeHandling, kernel utils.Kernel) *Filter {
	effect := new(Filter)
	effect.engine = f.engine
	effect.edgeHandling = edgeHandling
	effect.kernel = kernel

	return effect
}

func (f Factory) MultiFilter(edgeHandling strategy.EdgeHandling, resultMerging strategy.ColorMerger, kernels ...utils.Kernel) *MultiFilter {
	effect := new(MultiFilter)
	effect.engine = f.engine
	effect.edgeHandling = edgeHandling
	effect.resultMerging = resultMerging
	effect.kernels = kernels

	return effect
}

func (f Factory) Distortion(edgeHandling strategy.EdgeHandling, lens strategy.Lens) *Distortion {
	effect := new(Distortion)
	effect.engine = f.engine
	effect.edgeHandling = edgeHandling
	effect.lens = lens

	return effect
}

func (f Factory) Difference(diff strategy.ColorDifference) *Difference {
	effect := new(Difference)
	effect.engine = f.engine
	effect.diff = diff

	return effect
}

func (f Factory) ColorMapper(rules ...strategy.MappingRule) *ColorMapper {
	effect := new(ColorMapper)
	effect.engine = f.engine
	effect.rules = rules

	return effect
}

func (f Factory) Scale(colorSamplingStrategy strategy.ColorSampling) *Scale {
	effect := new(Scale)
	effect.engine = f.engine
	effect.colorSamplingStrategy = colorSamplingStrategy

	return effect
}

func (f Factory) Jitter(radius int) *Jitter {
	effect := new(Jitter)
	effect.engine = f.engine
	effect.radius = radius

	return effect
}

func (f Factory) Median(edgeHandling strategy.EdgeHandling, eval strategy.ColorEvaluation, windowRadius int) *Median {
	effect := new(Median)
	effect.engine = f.engine
	effect.edgeHandling = edgeHandling
	effect.eval = eval
	effect.windowRadius = windowRadius

	return effect
}
