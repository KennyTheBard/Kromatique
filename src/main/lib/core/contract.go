package core

import (
	"errors"
	"sync"
)

// Contract is a middle man object makes interface between the user
// and the execution engine, taking care of schedule in the process
type Contract interface {
	// PlaceOrder adds the given tasks in the executing schedule
	PlaceOrder(t Task) error
	// Deadline blocks the current routine execution until
	// the contract is fulfilled
	Deadline()
}

// PoolContract is an implementation of Contract that manages
// tasks that can be executed concurrently
type PoolContract struct {
	limit   int
	counter int
	wg      sync.WaitGroup
	q       chan TaskContract
}

func (s *PoolContract) PlaceOrder(t Task) error {
	if s.counter >= s.limit {
		return errors.New("contract limit has been reached")
	}

	s.counter++
	s.wg.Add(1)
	s.q <- TaskContract{task: t, wg: &s.wg}
	return nil
}

func (s *PoolContract) Deadline() {
	s.wg.Wait()
}

func NewPoolContract(orderSize int, q OrderQueue) *PoolContract {
	contract := new(PoolContract)
	contract.limit = orderSize
	contract.q = q

	return contract
}

// SequentialContract is an implementation of Contract that
// manages sequential execution of given task in the same goroutine
type SequentialContract struct {
	limit   int
	counter int
}

func (s *SequentialContract) PlaceOrder(t Task) error {
	if s.counter >= s.limit {
		return errors.New("contract limit has been reached")
	}

	s.counter++

	t()
	return nil
}

func (s *SequentialContract) Deadline() {
	return
}

func NewSequentialContract(orderSize int) *SequentialContract {
	contract := new(SequentialContract)
	contract.limit = orderSize

	return contract
}
