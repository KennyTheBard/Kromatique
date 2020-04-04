package distorsion

import (
	"../../utils"
)

type Lens interface {
	VecAt(x, y int) Vector
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

func (lens FishEyeLens) VecAt(x, y int) Vector {
	d := lens.center.Dist(utils.Pt2D(float64(x), float64(y)))
	if d > lens.radius || d == 0 {
		return VecZero()
	} else {
		alpha := d / lens.radius
		str := utils.LERP(0, lens.strength, alpha)

		beta := utils.Clamp(str/d, 0, 1)
		newX := utils.LERP(float64(x), lens.center.X, beta) - float64(x)
		newY := utils.LERP(float64(y), lens.center.Y, beta) - float64(y)

		return Vec(newX, newY)
	}
}
