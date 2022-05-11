package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}

	limitChan := make(chan struct{}, n)
	for _, task := range tasks {
		limitChan <- struct{}{}

		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				<-limitChan
			}()

			task()
		}()
	}

	wg.Wait()
	return nil
}
