package core

import (
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// Parallelize uses the current GOMAXPROCS to parallelize
// the given function for the requested number of times
func Parallelize(numLines int, fn func(y int)) {
	if numLines < 1 {
		return
	}

	procs := runtime.GOMAXPROCS(0)
	if procs > numLines {
		procs = numLines
	}

	c := make(chan int, numLines)
	for i := 0; i < numLines; i++ {
		c <- i
	}
	close(c)

	var wg sync.WaitGroup
	for i := 0; i < procs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range c {
				fn(i)
			}
		}()
	}
	wg.Wait()
}
