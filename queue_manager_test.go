package qumanager

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQueueManager_Up_Exit(t *testing.T) {
	queue := NewQueueManager(25)

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		queue.Up()
		wg.Add(1)

		go func(i int) {
			defer queue.Exit()
			defer wg.Done()

			time.Sleep(time.Second * time.Duration(1+rand.Intn(2)))
			fmt.Println("num:", i, "count:", queue.Count())
		}(i)
	}

	require.NotEqual(t, 0, queue.Count())
	wg.Wait()
	require.Equal(t, int64(0), queue.Count())
}

func TestTestQueueManager__Leak(t *testing.T) {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		counter     int64
	)

	go func() {
		qu := NewQueueManager(500)

		var wg sync.WaitGroup
		defer func() {
			wg.Wait()
			cancel()
		}()

		for i := 0; i < 2000; i++ {
			wg.Add(1)
			go func(i int) {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err)
					}
				}()

				qu.Up()

				defer qu.Exit()
				defer wg.Done()

				time.Sleep(time.Second * time.Duration(1+rand.Intn(5)))

				if i == 1000 {
					panic("panic -- OK")
				}

				atomic.AddInt64(&counter, 1)
			}(i)
		}
	}()

	time.Sleep(time.Second * 2)
	require.NotEmpty(t, atomic.LoadInt64(&counter))

wait:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("result:", atomic.LoadInt64(&counter))
			break wait
		case <-time.After(time.Second):
			fmt.Println(atomic.LoadInt64(&counter))
		}
	}

	require.Equal(t, int64(1999), atomic.LoadInt64(&counter))
	fmt.Println("ezz!")
}
