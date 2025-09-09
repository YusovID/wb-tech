package pool

import w "level-1/2/worker"

//
type Pool[Data any] struct {
	pool    chan *w.Worker
	workers []*w.Worker
	handler func(Data)
}

// Function New takes function handler
// and returns pointer on Pool struct
func New[Data any](handler func(Data)) *Pool[Data] {
	workers := w.New()

	return &Pool[Data]{
		pool:    make(chan *w.Worker, len(workers)),
		workers: workers,
		handler: handler,
	}
}

// Function Create fills up worker pool chan
func (p *Pool[Data]) Create() {
	for _, w := range p.workers {
		p.pool <- w
	}
}

// Function Handle takes Data, waits for new worker
// and starts function handler in goroutine
func (p *Pool[Data]) Handle(data Data) {
	w := <-p.pool

	go func() {
		p.handler(data)
		p.pool <- w
	}()
}

// Function Wait free worker pool channel
func (p *Pool[Data]) Wait() {
	for range p.workers {
		<-p.pool
	}
}
