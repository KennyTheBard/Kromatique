package plot

import "image"

type Plot interface {
	Render(image.Rectangle) image.Image
}
