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

func (q *QueueManager) Up() {
	q.queue <- struct{}{}
	atomic.AddInt64(&q.count, 1)
}

func (q *QueueManager) Exit() {
	<-q.queue
	atomic.AddInt64(&q.count, -1)
}

func (q *QueueManager) Count() int64 {
	return atomic.LoadInt64(&q.count)
}
