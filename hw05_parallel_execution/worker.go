package hw05parallelexecution

import (
	"fmt"
	"sync"
)

type Task func() error

type worker struct {
	id       int
	taskChan chan Task
}

func newWorker(id int, chanTask chan Task) *worker {
	return &worker{
		id:       id,
		taskChan: chanTask,
	}
}

func (w *worker) Start(wg *sync.WaitGroup, limiter *limiter) {
	fmt.Println("Start worker, id:", w.id)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range w.taskChan {
			if task() != nil {
				limiter.inc()
				if limiter.isLimitExceeded() {
					return
				}
			}
		}
	}()
}
