package main

import (
	"fmt"
	ac "level-1/18/atomicCounter"
	mc "level-1/18/mutexCounter"
	sc "level-1/18/simpleCounter"
	"sync"
)

// значение, на которое мы увеличиваем счетчики
const increment = 10000

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	// для ускорения работы программы запускаем "тесты" параллельно
	go testSimpleCounter(wg)
	go testMutexCounter(wg)
	go testAtomicCounter(wg)

	wg.Wait()
}

// функция testSimpleCounter симулирует полный цикл работы с обычным счетчиком
func testSimpleCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	simpleCounter := sc.New()
	simpleCounter.Increment(increment)

	fmt.Printf("value of simple counter: %d\n", simpleCounter.Value)
}

// функция testMutexCounter симулирует полный цикл работы со счетчиком, использующим mutex
func testMutexCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	mutexCounter := mc.New()
	mutexCounter.Increment(increment)

	fmt.Printf("value of mutex counter: %d\n", mutexCounter.Value)
}

// функция testAtomicCounter симулирует полный цикл работы со счетчиком, использующим atomic
func testAtomicCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	atomicCounter := ac.New()
	atomicCounter.Increment(increment)

	fmt.Printf("value of atomic counter: %d\n", atomicCounter.Value.Load())
}
