package services

type CounterService struct {
	count int
}

func NewCounterService() *CounterService {
	return &CounterService{
		count: 0,
	}
}

func (cs *CounterService) Increment() {
	cs.count++
}

func (cs *CounterService) Decrement() {
	cs.count--
}

func (cs *CounterService) Count() int {
	return cs.count
}
