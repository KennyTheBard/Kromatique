package core

import "sync"

// Task is a simple wrapper for a function that
// receives nothing and returns nothing
type Task func()

// TaskContract encapsulates a task and a reference
// to a WaitGroup in order to signalize completion
type TaskContract struct {
	task   Task
	taskWg *sync.WaitGroup
}

// OrderQueue is a short hand for a channel of TaskContract type
type OrderQueue chan TaskContract
