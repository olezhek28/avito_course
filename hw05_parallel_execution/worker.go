package hw05parallelexecution

import (
	"context"
	"fmt"
	"sync"
)

type Task func() error

type worker struct {
	id         int
	taskChan   chan Task
	resultChan chan error
}

func newWorker(id int, chanTask chan Task, resultChan chan error) *worker {
	return &worker{
		id:         id,
		taskChan:   chanTask,
		resultChan: resultChan,
	}
}

func (w *worker) Start(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("Start worker, id:", w.id)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range w.taskChan {
			w.resultChan <- task()
		}
	}()
}
