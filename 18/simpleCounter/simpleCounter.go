package simplecounter

import "sync"

// структура-обертка для простого типа int
type SimpleCounter struct {
	Value int
}

// функция New возвращает новый экземпляр структуры SimpleCounter
func New() *SimpleCounter {
	return &SimpleCounter{}
}

// функция Increment запускает increment горутин, которые увеличивают SimpleCounter.Value
func (sc *SimpleCounter) Increment(increment int) {
	wg := &sync.WaitGroup{}
	wg.Add(increment)

	for range increment {
		go func() {
			defer wg.Done()

			sc.Value++
		}()
	}

	wg.Wait()
}
