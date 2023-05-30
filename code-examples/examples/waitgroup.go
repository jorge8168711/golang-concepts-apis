package examples

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done() // every time the done is called, the wait group counter is decreased by 1
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}

func WgExample() {
	// the wait group  waits for a collection of goroutines to finish
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // this add 1 to the wait group counter
		go doSomething(i, &wg)
	}

	wg.Wait()
}
