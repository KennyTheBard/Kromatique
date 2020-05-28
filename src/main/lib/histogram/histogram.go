package histogram

import (
	"image"
	"image/color"

	"../core"
)

type IndexingRule interface {
	Color2Index(color.Color) int
	Min(int) color.Color
	Max(int) color.Color
}

type Histogram struct {
	engine core.Engine
	rule   IndexingRule
	values []int
}

func (histogram *Histogram) Scan(img image.Image) {
	//histogram.values[histogram.rule(c)] += 1
}
