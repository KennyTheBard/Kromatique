package strategy

import (
	"image"
	"math"

	"../geometry"
	"../utils"
)

// Vector is a struct than encapsulates an Euclidean vector that
// starts from origin of the space and is defined by its terminal point
type Vector struct {
	X, Y float64
}

func Vec(x, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

// VecZero is a constant vector that represents a point vector (magnitude 0)
var VecZero = Vec(0, 0)

// VectorMap encapsulates a map of vectors
type VectorMap struct {
	vs     [][]Vector
	bounds image.Rectangle
}

// At returns the vector at the given coordinates
func (m VectorMap) At(x, y int) Vector {
	return m.vs[y-m.bounds.Min.Y][x-m.bounds.Min.X]
}

// Set overrides the vector at the given coordinates with
// a given new vector
func (m *VectorMap) Set(x, y int, v Vector) {
	m.vs[y-m.bounds.Min.Y][x-m.bounds.Min.X] = v
}

// Bounds returns the subspace corresponding to the vectors
// contained in the map
func (m VectorMap) Bounds() image.Rectangle {
	return m.bounds
}

func NewVectorMap(bounds image.Rectangle) VectorMap {
	vs := make([][]Vector, bounds.Dy())
	for i := 0; i < bounds.Dy(); i++ {
		vs[i] = make([]Vector, bounds.Dx())
	}

	return VectorMap{vs: vs, bounds: bounds}
}

// Lens encapsulates the logic needed obtain distortion vectors for an image
type Lens interface {
	// VecAt returns the given distortion vector for a given position
	VecAt(x, y int) Vector
}

// PlanoLens encapsulates logic for a Plano lens (0.0 magnitude)
// that does not alter the image at all
type PlanoLens struct{}

func NewPlanoLens() *PlanoLens {
	l := new(PlanoLens)
	return l
}

func (lens PlanoLens) VecAt(x, y int) Vector {
	return VecZero
}

// LensOperation defines a general method of composing lens
type LensOperation func(Vector, Vector) Vector

// Add returns the addition of two vectors
func Add(a, b Vector) Vector {
	return Vec(a.X+b.X, a.Y+b.Y)
}

// Diff returns the subtraction of two vectors
func Diff(a, b Vector) Vector {
	return Vec(a.X-b.X, a.Y-b.Y)
}

// Merge returns the mean value of two vectors
func Merge(a, b Vector) Vector {
	return Vec((a.X+b.X)/2, (a.Y+b.Y)/2)
}

// LensAssembly encapsulates 2 Lenses and a LensOperation
// used to compose the two Lens
type LensAssembly struct {
	compose     LensOperation
	left, right Lens
}

func (asm LensAssembly) VecAt(x, y int) Vector {
	return asm.compose(asm.left.VecAt(x, y), asm.right.VecAt(x, y))
}

// FishEyeLens encapsulates logic for a fish eye distortion
type FishEyeLens struct {
	center   geometry.Point2D
	radius   float64
	strength float64
}

func NewFishEyeLens(center geometry.Point2D, radius, strength float64) *FishEyeLens {
	lens := new(FishEyeLens)
	lens.center = center
	lens.radius = radius
	lens.strength = strength

	return lens
}

func (lens FishEyeLens) VecAt(x, y int) Vector {
	d := lens.center.Dist(geometry.Pt2D(float64(x), float64(y)))
	if d > lens.radius || d == 0 {
		return VecZero
	} else {
		alpha := d / lens.radius
		str := utils.LERP(0, lens.strength, 1-utils.EaseInQuad(alpha))

		beta := utils.Clamp(str/d, 0, 1)
		newX := utils.LERP(float64(x), lens.center.X, beta) - float64(x)
		newY := utils.LERP(float64(y), lens.center.Y, beta) - float64(y)

		return Vec(newX, newY)
	}
}

// HorizontalWavesLens encapsulates logic for a wave effect
// on the horizontal direction
type HorizontalWavesLens struct {
	frequency, magnitude, translation float64
}

func NewHorizontalWavesLens(frequency, magnitude, translation float64) *HorizontalWavesLens {
	lens := new(HorizontalWavesLens)
	lens.frequency = frequency
	lens.magnitude = magnitude
	lens.translation = translation

	return lens
}

func (lens HorizontalWavesLens) VecAt(x, y int) Vector {
	return Vec(0, math.Sin((float64(x)+lens.translation)*lens.frequency)*lens.magnitude)
}

// RotationLens encapsulates logic for a fish eye distortion
type RotationLens struct {
	anchor  geometry.Point2D
	radians float64
}

func NewRotationLens(anchor geometry.Point2D, radians float64) *RotationLens {
	lens := new(RotationLens)
	lens.anchor = anchor
	lens.radians = radians

	return lens
}

func (lens RotationLens) VecAt(x, y int) Vector {
	realP := geometry.Pt2D(float64(x), float64(y))
	newP := geometry.Pt2D(realP.X, realP.Y)

	newP.Translate(-lens.anchor.X, -lens.anchor.Y)
	newP.Rotate(lens.radians)
	newP.Translate(lens.anchor.X, lens.anchor.Y)

	d := realP.Diff(newP)
	return Vec(d.X, d.Y)
}
