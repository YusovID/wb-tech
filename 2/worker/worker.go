package worker

var workersNum = 10

type Worker struct{}

func New() []*Worker {
	workers := make([]*Worker, 0, workersNum)

	for range workersNum {
		workers = append(workers, &Worker{})
	}

	return workers
}
