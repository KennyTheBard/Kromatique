package lib

type OrderQueue chan TaskContract

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
