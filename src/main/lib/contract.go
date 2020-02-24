package lib

import (
	"errors"
	"sync"
)

// OrderContract is a middle man object that serve as an interface
// between caller and engine in order to simplify the engine's job
type OrderContract struct {
	limit   int
	counter int
	wg      sync.WaitGroup
	q       chan TaskContract
}

// PlaceOrder adds a new order to the task queue,
// ensuring a method of getting feedback on the orders
func (s *OrderContract) PlaceOrder(t Task) error {
	if s.counter >= s.limit {
		return errors.New("contract limit has been reached")
	}

	s.counter++
	s.wg.Add(1)
	s.q <- TaskContract{task: t, wg: &s.wg}
	return nil
}

// Deadline should be called before any operation that depends on
// the result of the parallel operation is needed
func (s *OrderContract) Deadline() {
	s.wg.Wait()
}

func NewContract(orderSize int, q chan TaskContract) OrderContract {
	sup := OrderContract{}
	sup.limit = orderSize
	sup.q = q

	return sup
}