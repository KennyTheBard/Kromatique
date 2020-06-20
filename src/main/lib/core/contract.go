package core

import (
	"image"
	"sync"
)

// Contract is a middle man interface between the user and the execution engine,
// being responsible for scheduling in the process
type Contract interface {
	// PlaceOrder adds the given tasks in the executing schedule
	PlaceOrder(t Task)
	// Deadline blocks the current routine execution until
	// the contract is fulfilled
	Deadline()

	Promise(image.Image) *Promise
}

// PoolContract is an implementation of Contract interface that manages
// tasks that can be executed concurrently
type PoolContract struct {
	wg sync.WaitGroup
	q  chan TaskContract
}

func (s *PoolContract) PlaceOrder(t Task) {
	s.wg.Add(1)
	s.q <- TaskContract{task: t, taskWg: &s.wg}
}

func (s *PoolContract) Deadline() {
	s.wg.Wait()
}

func (s *PoolContract) Promise(img image.Image) *Promise {
	p := new(Promise)
	p.img = img
	p.contract = s

	return p
}

func NewPoolContract(q OrderQueue) *PoolContract {
	contract := new(PoolContract)
	contract.q = q

	return contract
}

// SequentialContract is an implementation of Contract interface that
// manages sequential execution of given task in the same goroutine
type SequentialContract struct{}

func (s *SequentialContract) PlaceOrder(t Task) {
	t()
}

func (s *SequentialContract) Deadline() {
	return
}

func (s *SequentialContract) Promise(img image.Image) *Promise {
	p := new(Promise)
	p.img = img
	p.contract = s

	return p
}

func NewSequentialContract() *SequentialContract {
	contract := new(SequentialContract)

	return contract
}
