package lib

import "sync"

type Task func()

type TaskContract struct {
	task Task
	wg   *sync.WaitGroup
}

