package core

import (
	"sync"
)

// Engine is an abstraction over the execution and schedule of tasks,
// encapsulating the logic and displaying a simple interface to its users
type Engine interface {
	// PoolContract returns a contract for the given number of tasks
	Contract() Contract
	// Stop closes the communication internal channel
	Stop()
}

// PoolEngine is the main engine designed for performance boost,
// using an internal pool of goroutines in order to execute given tasks
type PoolEngine struct {
	numWorkers int
	orderQueue OrderQueue
	engineWg   *sync.WaitGroup
}

func (engine PoolEngine) Contract() Contract {
	return NewPoolContract(engine.orderQueue)
}

func (engine PoolEngine) Stop() {
	close(engine.orderQueue)
	engine.engineWg.Wait()
}

func NewPoolEngine(numWorkers, queueSize int) *PoolEngine {
	engine := new(PoolEngine)
	engine.numWorkers = numWorkers
	engine.orderQueue = make(OrderQueue, queueSize)
	engine.engineWg = new(sync.WaitGroup)
	engine.engineWg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer engine.engineWg.Done()

			for tc := range engine.orderQueue {
				tc.task()
				tc.taskWg.Done()
			}
		}()
	}

	return engine
}

// SequentialEngine is an engine designed to function in entirely
// in the same goroutine as the invoking code
type SequentialEngine struct{}

func (engine SequentialEngine) Contract() Contract {
	return NewSequentialContract()
}

func (engine SequentialEngine) Stop() {}
