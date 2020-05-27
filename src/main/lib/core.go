package lib

import (
	"errors"
	"image"
	"sync"
)

// Promise encapsulates the image being worked on and a
// reference to the order contract of the appliance instance
type Promise struct {
	img      image.Image
	contract *OrderContract
}

// Result blocks the current goroutine until the image effect
// is applied completely or returns it immediately if it was
// already completed when is called
func (p Promise) Result() image.Image {
	(*p.contract).Deadline()
	return p.img
}

// NewPromise creates a new promise with the image still in work
// and a reference to the order contract of the appliance instance
func NewPromise(img image.Image, contract *OrderContract) *Promise {
	p := new(Promise)
	p.img = img
	p.contract = contract

	return p
}

// Task is a simple wrapper for a function that
// receives nothing and returns nothing
type Task func()

// TaskContract encapsulates a task and a reference
// to a WaitGroup in order to signalize completion
type TaskContract struct {
	task Task
	wg   *sync.WaitGroup
}

// OrderQueue is a short hand for a channel of TaskContract type
type OrderQueue chan TaskContract

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

// NewContract creates a new contract that accepts orderSize numbers if tasks
// and a communication channel with the engine in order to send tasks through
func NewContract(orderSize int, q OrderQueue) *OrderContract {
	sup := new(OrderContract)
	sup.limit = orderSize
	sup.q = q

	return sup
}

// KromEngine is the main parallelization structure of this library
// as every effect needs to encapsulate it in order to parallelize itself
type KromEngine struct {
	workForce  int
	orderQueue OrderQueue
}

// Contract returns a contract for the given number of tasks
func (krom *KromEngine) Contract(orderSize int) *OrderContract {
	return NewContract(orderSize, krom.orderQueue)
}

// Stop closes the communication channel with the internal workers
func (krom *KromEngine) Stop() {
	close(krom.orderQueue)
}

// NewKromEngine creates a new KromEngine with a given number of workers
// and a given task queue size
func NewKromEngine(workForce, queueSize int) *KromEngine {
	ke := new(KromEngine)
	ke.workForce = workForce
	ke.orderQueue = make(OrderQueue, queueSize)

	for i := 0; i < workForce; i++ {
		go func() {
			for {
				tc, more := <-ke.orderQueue
				if more {
					tc.task()
					tc.wg.Done()
				} else {
					break
				}
			}
		}()
	}

	return ke
}

// Base is a structure useful to ensure compatibility with the library
// as easy as possible, encapsulating interactions with the engine in use
type Base struct {
	engine *KromEngine
}

// GetEngine simple getter to access the engine currently in use
func (base *Base) GetEngine() *KromEngine {
	return base.engine
}
