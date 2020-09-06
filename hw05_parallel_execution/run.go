package hw05_parallel_execution // nolint:golint,stylecheck

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, n int, m int) error {
	var wg sync.WaitGroup
	var errCnt int64
	tCn := make(chan Task)
	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(ctxWithCancel, &wg, tCn, &errCnt)
	}

	var err error

	for _, task := range tasks {
		if m > 0 && atomic.LoadInt64(&errCnt) >= int64(m) {
			err = ErrErrorsLimitExceeded
			break
		}

		tCn <- task
	}

	cancelFunction()

	wg.Wait()

	return err
}

func worker(ctx context.Context, wg *sync.WaitGroup, tCn <-chan Task, errCnt *int64) {
	defer wg.Done()

	for {
		select {
		case t := <-tCn:
			if err := t(); err != nil {
				atomic.AddInt64(errCnt, 1)
			}
		case <-ctx.Done():
			return
		}
	}
}
