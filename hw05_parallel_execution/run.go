package hw05parallelexecution

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	workerPool := newPool(tasks, n, int64(m))
	workerPool.run()

	//wg := sync.WaitGroup{}

	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	////limitChan := make(chan struct{}, n)
	//errChan := make(chan error)
	//taskChan := make(chan Task, n)
	//
	//// Run workers
	//for i := 0; i < n; i++ {
	//	go worker(ctx, taskChan, errChan)
	//}
	//
	//// Processing tasks
	//for _, task := range tasks {
	//	taskChan <- task
	//}
	//
	//cancel()
	//close(taskChan)
	//
	//count := 0
	//select {
	//case err := <-errChan:
	//	if err != nil {
	//		count++
	//		if count >= m {
	//			cancel()
	//		}
	//	}
	//}

	//for _, task := range tasks {
	//	limitChan <- struct{}{}
	//
	//	wg.Add(1)
	//	go func() {
	//		var err error
	//
	//		defer func() {
	//			wg.Done()
	//			<-limitChan
	//			errChan <- err
	//		}()
	//
	//		err = task()
	//	}()
	//}
	//
	//wg.Wait()
	return nil
}

//func worker(ctx context.Context, taskChan chan Task, errChan chan error) {
//	select {
//	case task := <-taskChan:
//		errChan <- task()
//	case <-ctx.Done():
//		close(errChan)
//		return
//	}
//}

//func processing(errChan chan error) {
//
//	select {
//	case err := <-errChan:
//		if err != nil {
//
//		}
//	}
//}
