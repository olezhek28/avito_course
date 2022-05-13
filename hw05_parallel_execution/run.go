package hw05parallelexecution

import (
	"errors"
)

var (
	errErrorsLimitExceeded = errors.New("errors limit exceeded")
	errErrorsNoWorkers     = errors.New("no worker")
)

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return errErrorsNoWorkers
	}

	workerPool := newPool(tasks, n, int64(m))

	return workerPool.run()
}
