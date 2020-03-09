package render

import (
	core ".."
	"../geometry"
	"../utils"
	"fmt"
	"image"
	"image/draw"
	"math"
)

type Line struct {
	core.Base
	path     geometry.Path
	renderer Renderer
}

func (l *Line) Render(bounds image.Rectangle) *core.Promise {
	// create a proximity map in order to save how far is each
	// pixel from the path used to render the line
	proximityMap := make([][]int, bounds.Dy())
	for i := 0; i < bounds.Dy(); i++ {
		proximityMap[i] = make([]int, bounds.Dx())
	}

	// create a queue of closest pixels to the path
	queue := make([]image.Point, 0)
	for i := 0.0; i < 1.0; i += 0.01 {
		// approximate the pixel coordinates
		p := l.path.GetPoint(i)
		y := int(math.Floor(p.Y))
		x := int(math.Floor(p.X))

		// reduce redundancy in the queue loop
		if proximityMap[y][x] == 0 {
			queue = append(queue, image.Pt(x, y))
		}

		// set base value for the closest pixels to the paths
		proximityMap[y][x] = 1
	}

	// iterate through the queued pixels
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		// obtain Von Neumann neighbours
		neighbours := make([]image.Point, 0)
		if inRectangle(p.X, p.Y-1, bounds) {
			neighbours = append(neighbours, image.Pt(p.X, p.Y-1))
		}
		if inRectangle(p.X-1, p.Y, bounds) {
			neighbours = append(neighbours, image.Pt(p.X-1, p.Y))
		}
		if inRectangle(p.X+1, p.Y, bounds) {
			neighbours = append(neighbours, image.Pt(p.X+1, p.Y))
		}
		if inRectangle(p.X, p.Y+1, bounds) {
			neighbours = append(neighbours, image.Pt(p.X, p.Y+1))
		}

		// compute proximity value for current pixel based on
		// highest value among the neighbours
		if proximityMap[p.Y][p.X] == 0 {
			max := 0
			for _, n := range neighbours {
				if max < proximityMap[n.Y][n.X] {
					max = proximityMap[n.Y][n.X]
				}
			}

			proximityMap[p.Y][p.X] = max + 1
		}

		// when the limit width has been reached, there won't
		// be any new neighbours added in the queue
		if proximityMap[p.Y][p.X] == l.renderer.Width() {
			continue
		}

		// add unexplored neighbouring pixels in the queue
		for _, n := range neighbours {
			if proximityMap[n.Y][n.X] == 0 {
				queue = append(queue, n)
			}
		}
	}

	ret := utils.CreateRGBA(bounds)
	contract := l.GetEngine().Contract(bounds.Dy())

	for i := bounds.Min.Y; i < bounds.Max.Y; i++ {
		y := i
		if err := contract.PlaceOrder(func() {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				ret.(draw.Image).Set(x, y, l.renderer.Render(proximityMap[y][x]))
			}
		}); err != nil {
			fmt.Print(err)
			break
		}
	}

	return core.NewPromise(ret, contract)
}

func inRectangle(x, y int, rect image.Rectangle) bool {
	return x >= rect.Min.X && x <= rect.Max.X && y >= rect.Min.Y && y <= rect.Max.Y
}

func NewLine(path geometry.Path, renderer Renderer) *Line {
	l := new(Line)
	l.path = path
	l.renderer = renderer

	return l
}
