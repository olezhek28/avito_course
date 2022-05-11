package hw05parallelexecution

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	limitChan := make(chan struct{}, n)
	for _, t := range tasks {
		limitChan <- struct{}{}

		go t()
	}

	return nil
}
