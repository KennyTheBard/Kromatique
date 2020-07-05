package plot

import "image"

// Plot is an interface for any king of plot
type Plot interface {
	// Render returns an image of the given bounds
	// displaying the encapsulated data
	Render(image.Rectangle) image.Image
}
