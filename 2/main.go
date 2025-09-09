package main

import (
	"fmt"
	workerPool "level-1/2/pool"
)

var nums = []int{2, 4, 6, 8, 10}

type IPool interface {
	Create()
	Handle(int)
	Wait()
}

func main() {
	var pool IPool = workerPool.New(square)

	// fill up worker pool channel
	pool.Create()

	for _, num := range nums {
		// running hanler function
		pool.Handle(num)
	}

	// waiting for all goroutines
	pool.Wait()
}

func square(num int) {
	fmt.Println(num * num)
}
