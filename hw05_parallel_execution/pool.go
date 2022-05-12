package hw05parallelexecution

import (
	"context"
	"sync"
)

const chanSize = 1000

type pool struct {
	tasks         []Task
	concurrency   int
	collector     chan Task
	wg            sync.WaitGroup
	maxErrorCount int
	errorCount    int
	resultChan    chan error
	mt            sync.Mutex
}

func newPool(tasks []Task, concurrency int, maxErrorCount int) *pool {
	return &pool{
		tasks:         tasks,
		concurrency:   concurrency,
		collector:     make(chan Task, chanSize),
		maxErrorCount: maxErrorCount,
		resultChan:    make(chan error, chanSize),
	}
}

func (p *pool) run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	// Run workers
	for i := 0; i < p.concurrency; i++ {
		w := newWorker(i, p.collector, p.resultChan)
		w.Start(ctx, &p.wg)
	}

	// Processing tasks
	for _, task := range p.tasks {
		p.collector <- task
	}
	close(p.collector)

	for err := range p.resultChan {
		if err != nil {
			p.errorCount++
		}

		if p.errorCount >= p.maxErrorCount {
			cancel()
		}
	}

	p.wg.Wait()
}

//func (p *pool) getErrorCount() int {
//	p.mt.Lock()
//	defer p.mt.Unlock()
//
//	return p.errorCount
//}
