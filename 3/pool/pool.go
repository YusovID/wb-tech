// Package pool implements a generic worker pool.
// It manages a collection of workers to process data concurrently.
package pool

import (
	"fmt"
	// w is an alias for the worker package.
	w "level-1/3/worker"
)

// Pool represents a generic worker pool.
// It is responsible for managing a set of workers and distributing data among them.
type Pool[Data any] struct {
	// workers is a slice containing all worker instances managed by the pool.
	workers []*w.Worker
	// pool is a channel that acts as a queue of available workers.
	pool chan *w.Worker
	// handler is the function that will be executed by a worker to process data.
	handler func(int, Data)
}

// New creates and returns a new generic worker pool.
// It initializes the workers and the pool channel.
// The 'handler' parameter is a function that defines the work to be done.
func New[Data any](handler func(int, Data)) *Pool[Data] {
	// Create the initial set of workers.
	workers := w.New()

	// Return a new Pool instance.
	return &Pool[Data]{
		workers: workers,
		// The pool channel is buffered with a size equal to the number of workers.
		pool:    make(chan *w.Worker, len(workers)),
		handler: handler,
	}
}

// Create makes all workers available for processing by sending them to the pool channel.
// This should be called before handling a new batch of data.
func (p *Pool[Data]) Create() {
	for _, w := range p.workers {
		p.pool <- w
	}
}

// Handle takes a piece of data and assigns it to an available worker.
// It retrieves a worker from the pool, starts a new goroutine for processing,
// and returns the worker to the pool once the processing is complete.
func (p *Pool[Data]) Handle(data Data) {
	// Wait for and get an available worker from the pool.
	w := <-p.pool

	// Start a new goroutine to process the data.
	go func() {
		// Execute the handler function with the worker's ID and the data.
		p.handler(w.Id, data)
		// Increment the worker's processed data counter.
		w.ProcessedDataCount++
		// Return the worker to the pool, making it available for another task.
		p.pool <- w
	}()
}

// Wait blocks until all workers have finished their tasks and returned to the pool.
// This is used to synchronize after a batch of data has been dispatched.
func (p *Pool[Data]) Wait() {
	for range p.workers {
		<-p.pool
	}
}

// Stats prints a summary of the work done by each worker.
func (p *Pool[Data]) Stats() {
	fmt.Println("-----------------------STATISTICS-------------------------")
	for _, w := range p.workers {
		fmt.Printf("worker '%d' processed %d datas\n", w.Id, w.ProcessedDataCount)
	}
}
