package core

import (
	"sync"
)

// Engine is an abstraction over the execution and schedule of tasks,
// encapsulating the logic and displaying a simple interface to its users
type Engine interface {
	// PlaceOrder adds the given tasks in the executing schedule
	PlaceOrder(...Task)
	// Stop closes the communication internal channel
	Stop()
}

// PoolEngine is the main engine designed for performance boost,
// using an internal pool of goroutines in order to execute given tasks
type PoolEngine struct {
	orderQueue chan Task
	engineWg   *sync.WaitGroup
}

func (engine *PoolEngine) PlaceOrder(tasks ...Task) {
	for _, t := range tasks {
		engine.orderQueue <- t
	}
}

func (engine PoolEngine) Stop() {
	close(engine.orderQueue)
	engine.engineWg.Wait()
}

func NewPoolEngine(numWorkers, queueSize int) *PoolEngine {
	engine := new(PoolEngine)
	engine.orderQueue = make(chan Task, queueSize)
	engine.engineWg = new(sync.WaitGroup)
	engine.engineWg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer engine.engineWg.Done()

			for task := range engine.orderQueue {
				task()
			}
		}()
	}

	return engine
}

// SequentialEngine is an engine designed to function in entirely
// in the same goroutine as the invoking code
type SequentialEngine struct{}

func (engine *SequentialEngine) PlaceOrder(tasks ...Task) {
	for _, t := range tasks {
		t()
	}
}

func (engine *SequentialEngine) Stop() {}

// Task is a simple wrapper for a function that
// receives nothing and returns nothing
type Task func()
