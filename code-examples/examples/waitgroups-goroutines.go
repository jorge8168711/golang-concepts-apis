// CONCURRENCY

package examples

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency is dealing multiple things at a single time
// while parallelism is doing multiple things at the same time.

// A goroutine is a lightweight thread of execution.
func say(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func GoRoutineEg() {
	// this runs synchronously
	say("direct")

	// to invoke this function in a goroutine, use go say(s)
	go say("goroutine")

	// it is possible to start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// our two function calls are running asynchronously in separate goroutines now.
	// wait for them to finish, a better approach is to use a WaitGroup
	time.Sleep(time.Second) // this is a bad practice
	fmt.Println("done")
}

// To wait for multiple goroutines to finish, we can use a wait group.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// Note: if a WaitGroup is explicitly passed into
// functions, it should be done by pointer
func workerWg(id int, wg *sync.WaitGroup) {
	defer wg.Done() // with this we release the goroutine from the WaitGroup
	fmt.Printf("Worker %d starting\n", id)

	// sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroupEg() {
	// the WaitGroup is used to wait for all the goroutines launched here to finish.
	var wg sync.WaitGroup

	// launch several goroutines and increment the WaitGroup counter for each
	wg.Add(1)

	// we wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
	// This way the worker itself does not have to be aware of the concurrency primitives
	// involved in its execution.
	go func(msg string) {
		defer wg.Done()
		worker(1)
	}("going")

	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		worker(6)
	}("other")

	wg.Add(1)
	workerWg(8, &wg)

	wg.Wait()
}
