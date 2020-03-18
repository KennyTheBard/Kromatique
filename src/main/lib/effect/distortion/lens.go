package distorsion

import (
	"../../utils"
	"image"
	"math"
)

type Lens interface {
	VectorMap() VectorMap
}

type LensAssembly Lens

type FishEyeLens struct {
	center   utils.Point2D
	radius   float64
	strength float64
}

func NewFishEyeLens(center utils.Point2D, radius, strength float64) *FishEyeLens {
	fel := new(FishEyeLens)
	fel.center = center
	fel.radius = radius
	fel.strength = strength

	return fel
}

func (lens FishEyeLens) VectorMap() VectorMap {
	bounds := image.Rect(
		int(math.Floor(lens.center.X-lens.radius)),
		int(math.Floor(lens.center.Y-lens.radius)),
		int(math.Ceil(lens.center.X+lens.radius)),
		int(math.Ceil(lens.center.Y+lens.radius)))
	vm := NewVectorMap(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			d := lens.center.Dist(utils.Pt2D(float64(x), float64(y)))
			if d > lens.radius || d == 0 {
				vm.Set(x, y, VecZero())
			} else {
				alpha := d / lens.radius
				str := utils.LERP(0, lens.strength, alpha)

				beta := utils.Clamp(str/d, 0, 1)
				newX := utils.LERP(float64(x), lens.center.X, beta) - float64(x)
				newY := utils.LERP(float64(y), lens.center.Y, beta) - float64(y)

				vm.Set(x, y, Vec(newX, newY))
			}
		}
	}

	return vm
}
