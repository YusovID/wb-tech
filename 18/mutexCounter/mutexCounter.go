package mutexcounter

import "sync"

// структура-обертка для типа int и mutex
type MutexCounter struct {
	Value int64
	mu    *sync.Mutex
}

// функция New возвращает экземпляр структуры MutexCounter
func New() *MutexCounter {
	return &MutexCounter{
		mu: &sync.Mutex{},
	}
}

// функция Increment запускает increment горутин, которые увеличивают MutexCounter.Value
func (mc *MutexCounter) Increment(increment int) {
	wg := &sync.WaitGroup{}
	wg.Add(increment)

	for range increment {
		go func() {
			defer wg.Done()

			mc.mu.Lock()
			mc.Value++
			mc.mu.Unlock()
		}()
	}

	wg.Wait()
}
