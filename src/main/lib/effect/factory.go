package effect

import (
	"../core"
	"../utils"
	"image"
)

type Factory struct {
	engine core.Engine
}

func NewFactory(engine core.Engine) *Factory {
	f := new(Factory)
	f.engine = engine

	return f
}

func (f Factory) Flip(strategy FlipStrategy) *Flip {
	effect := new(Flip)
	effect.engine = f.engine
	effect.strategy = strategy

	return effect
}

func (f Factory) SingleKernelFilter(edgeHandling EdgeHandlingStrategy, kernel utils.Matrix) *SingleKernel {
	effect := new(SingleKernel)
	effect.engine = f.engine
	effect.edgeHandling = edgeHandling
	effect.kernel = kernel

	return effect
}

func (f Factory) MultiKernelFilter(edgeHandling EdgeHandlingStrategy, resultMerging ResultMergingStrategy, kernels []utils.Matrix) *MultiKernel {
	effect := new(MultiKernel)
	effect.engine = f.engine
	effect.edgeHandling = edgeHandling
	effect.resultMerging = resultMerging
	effect.kernels = kernels

	return effect
}

func (f Factory) Distortion(edgeHandling EdgeHandlingStrategy, lens Lens) *Distortion {
	effect := new(Distortion)
	effect.engine = f.engine
	effect.edgeHandling = edgeHandling
	effect.lens = lens

	return effect
}

func (f Factory) Difference(diff DifferenceStrategy) *Difference {
	effect := new(Difference)
	effect.engine = f.engine
	effect.diff = diff

	return effect
}

func (f Factory) ColorMapper(rules []MappingRule) *ColorMapper {
	effect := new(ColorMapper)
	effect.engine = f.engine
	effect.rules = rules

	return effect
}

func (f Factory) Normalization(source, target *utils.ColorInterval) *Normalization {
	effect := new(Normalization)
	effect.engine = f.engine
	effect.sourceInterval = source
	effect.targetInterval = target

	return effect
}

func (f Factory) Overlay(stamp image.Image, origin image.Point, opacity float64) *Overlay {
	effect := new(Overlay)
	effect.engine = f.engine
	effect.stamp = stamp
	effect.origin = origin
	effect.opacity = opacity

	return effect
}

func (f Factory) Scale(scaleFactorStrategy ScaleFactorStrategy, colorSamplingStrategy ColorSamplingStrategy) *Scale {
	effect := new(Scale)
	effect.engine = f.engine
	effect.scaleFactorStrategy = scaleFactorStrategy
	effect.colorSamplingStrategy = colorSamplingStrategy

	return effect
}
