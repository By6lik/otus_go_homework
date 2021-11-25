package hw05parallelexecution

import (
	"errors"
	"sync/atomic"
	// "time"
	// "math/rand"
	// "sync/atomic"
	"sync"
)

var ( 
ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
ErrInlvalidn = errors.New("n should be more than zero")
)

type Task func() error


// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrInlvalidn
	}
	if m <= 0 {
		m = len(tasks) + 1
	}
	wg := sync.WaitGroup{}
	ch := make(chan Task)
	var errCount int32
	wg.Add(n)
	for i := 0; i < n; i++ {
		//consumer
		go func() {
			defer wg.Done() 
			for task := range ch {
				if err := task(); err!=nil {
					atomic.AddInt32(&errCount, 1)
				}
			}
		}()
		}
		//producer
		for _, task := range tasks {
			if  atomic.LoadInt32(&errCount)>= int32(m) {
				break
			}
			ch <- task
		}
	close(ch)
	wg.Wait()
	if  errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}