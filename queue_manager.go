package qumanager

import (
	"sync/atomic"
)

type QueueManager struct {
	queue chan struct{}
	count int64
}

func NewQueueManager(size int64) *QueueManager {
	return &QueueManager{
		queue: make(chan struct{}, size),
	}
}

// entering the queue
// If the place is occupied, the function will wait until someone leaves the queue
func (q *QueueManager) Up() {
	q.queue <- struct{}{}
	atomic.AddInt64(&q.count, 1)
}

// getting out of the queue
func (q *QueueManager) Exit() {
	<-q.queue
	atomic.AddInt64(&q.count, -1)
}

// current queue count
func (q *QueueManager) Count() int64 {
	return atomic.LoadInt64(&q.count)
}
