package lib

import "image"

type Promise struct {
	img image.Image
	contract *OrderContract
}

func (p *Promise) Result() image.Image {
	(*p.contract).Deadline()
	return p.img
}

func NewPromise(img image.Image, contract *OrderContract) Promise {
	return Promise{img:img, contract:contract}
}
