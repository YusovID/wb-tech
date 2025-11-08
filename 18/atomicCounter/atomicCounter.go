package atomiccounter

import (
	"sync"
	"sync/atomic"
)

// структура-обертка для создания методов у atomic
type AtomicCounter struct {
	Value atomic.Int64
}

// функция New создает новый экземпляр структуры AtomicCounter
func New() *AtomicCounter {
	return &AtomicCounter{
		Value: atomic.Int64{},
	}
}

// функция Increment запускает increment горутин, которые увеличивают AtomicCounter.Value
func (ac *AtomicCounter) Increment(increment int) {
	wg := &sync.WaitGroup{}
	wg.Add(increment)

	for range increment {
		go func() {
			defer wg.Done()

			ac.Value.Add(1)
		}()
	}

	wg.Wait()
}
