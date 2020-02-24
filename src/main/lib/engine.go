package lib

type KromEngine struct {
	workForce  int
	orderQueue chan TaskContract
}

func (krom *KromEngine) Contract(orderSize int) OrderContract {
	return NewContract(orderSize, krom.orderQueue)
}

func (krom *KromEngine) Stop() {
	close(krom.orderQueue)
}

func NewKromEngine(workForce, queueSize int) KromEngine {
	ke := KromEngine{workForce: workForce}
	ke.orderQueue = make(chan TaskContract, queueSize)

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