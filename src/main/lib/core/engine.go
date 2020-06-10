package core

import (
	"sync"
)

// Engine is an abstraction over the execution and schedule of tasks,
// encapsulating the logic and displaying a simple interface to its users
type Engine interface {
	// PoolContract returns a contract for the given number of tasks
	Contract(int) Contract
	// Stop closes the communication internal channel
	Stop()
}

// PoolEngine is the main engine designed for performance boost,
// using an internal pool of goroutines in order to execute given tasks
type PoolEngine struct {
	workForce  int
	orderQueue OrderQueue
	wg         *sync.WaitGroup
}

func (engine PoolEngine) Contract(orderSize int) Contract {
	return NewPoolContract(orderSize, engine.orderQueue)
}

func (engine PoolEngine) Stop() {
	close(engine.orderQueue)
	engine.wg.Wait()
}

func NewPoolEngine(workForce, queueSize int) *PoolEngine {
	ke := new(PoolEngine)
	ke.workForce = workForce
	ke.orderQueue = make(OrderQueue, queueSize)
	ke.wg = new(sync.WaitGroup)
	ke.wg.Add(workForce)

	for i := 0; i < workForce; i++ {
		go func() {
			defer ke.wg.Done()

			for tc := range ke.orderQueue {
				tc.task()
				tc.wg.Done()
			}
		}()
	}

	return ke
}

// SequentialEngine is an engine designed to function in entirely
// in the same goroutine as the invoking code
type SequentialEngine struct{}

func (engine SequentialEngine) Contract(orderSize int) Contract {
	return NewSequentialContract(orderSize)
}

func (engine SequentialEngine) Stop() {}
