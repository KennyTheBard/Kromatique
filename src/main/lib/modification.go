package lib

import "image"

// ImageModification is the core interface of all image processing effects
type ImageModification interface {
	Apply(image.Image) Promise
}