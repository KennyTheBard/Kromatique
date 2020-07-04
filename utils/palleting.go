package utils

import (
	"image"
	"image/color"
	"math"
)

type interval struct {
	start, end int
}

func buildInterval(start, end int) interval {
	return interval{
		start: start,
		end:   end,
	}
}

func (i interval) halfInterval(first bool) interval {
	middle := (i.start + i.end) / 2
	if first {
		return buildInterval(i.start, middle)
	} else {
		return buildInterval(middle+1, i.end)
	}
}

type octree interface {
	countLeafs() int
	minimumLeafSize() int
	prune(int) bool
	getColors() []color.Color
}

type branch struct {
	subs []octree
}

func (b *branch) countLeafs() int {
	total := 0
	for _, s := range b.subs {
		total += s.countLeafs()
	}
	return total
}

func (b *branch) minimumLeafSize() int {
	min := math.MaxInt32
	for _, s := range b.subs {
		if currMin := s.minimumLeafSize(); currMin < min {
			min = currMin
		}
	}
	return min
}

func (b *branch) prune(th int) bool {
	removeIdx := make([]int, 0)
	for idx, s := range b.subs {
		if s.prune(th) {
			removeIdx = append(removeIdx, idx)
		}
	}

	aux := make([]octree, 0)
mainLoop:
	for idx, s := range b.subs {
		for _, i := range removeIdx {
			if i == idx {
				continue mainLoop
			}
		}

		aux = append(aux, s)
	}

	b.subs = aux
	return len(b.subs) == 0
}

func (b *branch) getColors() []color.Color {
	cols := make([]color.Color, 0)
	for _, s := range b.subs {
		cols = append(cols, s.getColors()...)
	}
	return cols
}

type leaf struct {
	ix, iy, iz interval
	colorCube  *[][][]int
	pixels     int
}

func (l *leaf) countLeafs() int {
	return 1
}

func (l *leaf) minimumLeafSize() int {
	return l.pixels
}

func (l *leaf) prune(th int) bool {
	return l.pixels <= th
}

func (l *leaf) getColors() []color.Color {
	var tr, tg, tb int
	for x := l.ix.start; x < l.ix.end; x++ {
		for y := l.iy.start; y < l.iy.end; y++ {
			for z := l.iz.start; z < l.iz.end; z++ {
				tr += (*l.colorCube)[x][y][z] * x
				tg += (*l.colorCube)[x][y][z] * y
				tb += (*l.colorCube)[x][y][z] * z
			}
		}
	}

	return []color.Color{
		color.RGBA{
			R: uint8(math.Round(float64(tr) / float64(l.pixels))),
			G: uint8(math.Round(float64(tg) / float64(l.pixels))),
			B: uint8(math.Round(float64(tb) / float64(l.pixels))),
			A: math.MaxUint8,
		},
	}
}

func (l *leaf) divideLeaf(firstX, firstY, firstZ bool) leaf {
	ol := leaf{
		ix:        l.ix.halfInterval(firstX),
		iy:        l.iy.halfInterval(firstY),
		iz:        l.iz.halfInterval(firstZ),
		colorCube: l.colorCube,
	}
	for x := ol.ix.start; x < ol.ix.end; x++ {
		for y := ol.iy.start; y < ol.iy.end; y++ {
			for z := ol.iz.start; z < ol.iz.end; z++ {
				ol.pixels += (*ol.colorCube)[x][y][z]
			}
		}
	}
	return ol
}

func buildTree(tree leaf, th int) octree {
	total := 0
	for x := tree.ix.start; x < tree.ix.end; x++ {
		for y := tree.iy.start; y < tree.iy.end; y++ {
			for z := tree.iz.start; z < tree.iz.end; z++ {
				total += (*tree.colorCube)[x][y][z]
			}
		}

		if total > th && tree.ix.start != tree.ix.end {
			return &branch{subs: []octree{
				buildTree(tree.divideLeaf(false, false, false), th),
				buildTree(tree.divideLeaf(false, false, true), th),
				buildTree(tree.divideLeaf(false, true, false), th),
				buildTree(tree.divideLeaf(false, true, true), th),
				buildTree(tree.divideLeaf(true, false, false), th),
				buildTree(tree.divideLeaf(true, false, true), th),
				buildTree(tree.divideLeaf(true, true, false), th),
				buildTree(tree.divideLeaf(true, true, true), th),
			}}
		}
	}

	return &tree
}

func GeneratePallet(img image.Image, numColors int) []color.Color {
	// initialize color cube
	colorCube := make([][][]int, 256)
	for i := 0; i < 256; i++ {
		colorCube[i] = make([][]int, 256)
		for j := 0; j < 256; j++ {
			colorCube[i][j] = make([]int, 256)
		}
	}

	// add each color in the cube
	countColor := 0
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r, g, b = r>>8, g>>8, b>>8

			if colorCube[r][g][b] == 0 {
				countColor += 1
			}
			colorCube[r][g][b] += 1
		}
	}

	// build complete tree
	th := int(math.Ceil(float64(countColor) / float64(numColors)))
	tree := buildTree(leaf{
		ix:        interval{0, 256},
		iy:        interval{0, 256},
		iz:        interval{0, 256},
		colorCube: &colorCube}, th)

	for tree.countLeafs() > numColors {
		tree.prune(tree.minimumLeafSize())
	}

	return tree.getColors()
}
