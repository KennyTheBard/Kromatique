package lib

import (
	"errors"
	"sync"
)

type Task func()

type taskContract struct {
	task Task
	wg   *sync.WaitGroup
}

// OrderContract is a middle man object that serve as an interface
// between caller and engine in order to simplify the engine's job
type OrderContract interface {
	// PlaceOrder adds a new order to the task queue,
	// ensuring a method of getting feedback on the orders
	PlaceOrder(Task) error
	// Deadline should be called before any operation that depends on
	// the result of the parallel operation is needed
	Deadline()
}

type contract struct {
	limit   int
	counter int
	wg      sync.WaitGroup
	q       chan taskContract
}

func (s *contract) PlaceOrder(t Task) error {
	if s.counter >= s.limit {
		return errors.New("contract limit has been reached")
	}

	s.counter++
	s.wg.Add(1)
	s.q <- taskContract{task: t, wg: &s.wg}
	return nil
}

func (s *contract) Deadline() {
	s.wg.Wait()
}

func createContract(orderSize int, q chan taskContract) OrderContract {
	sup := contract{}
	sup.limit = orderSize
	sup.q = q

	return &sup
}

type KromEngine struct {
	workForce  int
	orderQueue chan taskContract
}

func (krom *KromEngine) Contract(orderSize int) OrderContract {
	return createContract(orderSize, krom.orderQueue)
}

func (krom *KromEngine) Stop() {
	close(krom.orderQueue)
}

func NewKromEngine(workForce, queueSize int) KromEngine {
	ke := KromEngine{workForce: workForce}
	ke.orderQueue = make(chan taskContract, queueSize)

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

func (krom *KromEngine) Grayscale() *Grayscale {
	return &Grayscale{engine:krom}
}

func (krom *KromEngine) CustomGrayscale(redRatio, greenRatio, blueRatio int) *Grayscale {
	return &Grayscale{redRatio: redRatio, greenRatio: greenRatio, blueRatio: blueRatio, engine:krom}
}

func (krom *KromEngine) Sepia() *Sepia {
	return &Sepia{engine:krom}
}