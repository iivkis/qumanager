# Queue Manager | Description
---
This package allows you to limit the number of simultaneously executed goroutines and ensures that no more than N goroutines will be executed at the same time.
# How to use?

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/iivkis/qumanager"
)

func main() {
	var (
		// creating a new queue of size 10
		qu = qumanager.NewQueueManager(10)

		// it is necessary to prevent premature termination of the program
		wg = sync.WaitGroup{}
	)

	// we expect the completion of all goroutines at the end of the program
	defer wg.Wait()

	// we say that we need to wait for the end of 30 goroutines
	wg.Add(30)

	for index := 0; index < 30; index++ {
		// we take a place in the queue. If the place is occupied,
		// the function will wait until someone leaves the queue
		qu.Up()

		go func(index int) {
			//we say that goroutine has been completed
			defer wg.Done()

			// getting out of the queue
			defer qu.Exit()

			// print the information & sleep
			fmt.Printf("index: %d; count: %d\n", index, qu.Count())
			time.Sleep(time.Second)
		}(index)
	}

	/* Output:
	index: 8; count: 10
	index: 0; count: 10
	index: 2; count: 10
	index: 3; count: 10
	index: 4; count: 10
	index: 5; count: 10
	index: 6; count: 10
	index: 7; count: 10
	index: 1; count: 7
	index: 9; count: 10
	index: 10; count: 4
	index: 11; count: 6
	index: 13; count: 6
	index: 12; count: 6
	index: 14; count: 7
	index: 16; count: 9
	index: 18; count: 10
	index: 15; count: 8
	index: 19; count: 10
	index: 17; count: 10
	index: 20; count: 10
	index: 22; count: 10
	index: 21; count: 10
	index: 23; count: 10
	index: 29; count: 10
	index: 24; count: 10
	index: 27; count: 10
	index: 28; count: 10
	index: 26; count: 10
	index: 25; count: 10
	*/
}

```
