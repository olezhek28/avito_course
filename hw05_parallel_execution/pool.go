package hw05parallelexecution

import (
	"sync"
	"sync/atomic"
)

const chanSize = 1000

type limiter struct {
	count int64
	limit int64
}

func newLimiter(limit int64) *limiter {
	return &limiter{
		limit: limit,
	}
}

func (l *limiter) inc() {
	atomic.AddInt64(&l.count, 1)
}

func (l *limiter) isLimitExceeded() bool {
	return atomic.LoadInt64(&l.count) >= atomic.LoadInt64(&l.limit)
}

type pool struct {
	tasks       []Task
	concurrency int
	collector   chan Task
	wg          sync.WaitGroup
	limiter     *limiter
}

func newPool(tasks []Task, concurrency int, maxErrorCount int64) *pool {
	return &pool{
		tasks:       tasks,
		concurrency: concurrency,
		collector:   make(chan Task, chanSize),
		limiter:     newLimiter(maxErrorCount),
	}
}

func (p *pool) run() error {
	// Run workers
	for i := 0; i < p.concurrency; i++ {
		w := newWorker(i, p.collector)
		w.Start(&p.wg, p.limiter)
	}

	// Processing tasks
	for _, task := range p.tasks {
		p.collector <- task
	}
	close(p.collector)

	p.wg.Wait()

	if p.limiter.isLimitExceeded() {
		return ErrErrorsLimitExceeded
	}

	return nil
}
