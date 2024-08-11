package timeoutwrapper

import (
	"sync"
	"sync/atomic"
)

// Custom struct to keep track of active workers
type WorkerGroup struct {
	wg      sync.WaitGroup
	counter int32 // Use int32 for atomic operations
}

func (wg *WorkerGroup) Add(delta int) {
	atomic.AddInt32(&wg.counter, int32(delta))
	wg.wg.Add(delta)
}

func (wg *WorkerGroup) Done() {
	atomic.AddInt32(&wg.counter, -1)
	wg.wg.Done()
}

func (wg *WorkerGroup) Wait() {
	wg.wg.Wait()
}

func (wg *WorkerGroup) ActiveWorkers() int32 {
	return atomic.LoadInt32(&wg.counter)
}

func (wg *WorkerGroup) FinishAllWorkers() {
	ctr := wg.counter
	for i := 0; i < int(ctr); i++ {
		wg.Done()
	}
}
