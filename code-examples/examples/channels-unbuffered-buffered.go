package examples

import (
	"fmt"
	"sync"
	"time"
)

// By default channels are unbuffered, meaning that they will only
// accept sends (chan <-) if there is a corresponding
// receive (<- chan) ready to receive the sent value.
func UnBuffChannels() {
	// a channel is always blocking if is full
	// and an unbuffered channel is always full, because the default size is 0
	c := make(chan int)

	// the way to fix this is to use a goroutine
	// with this code the result variable is ready to receive the channel value
	// with this the channel still blocking, but doesn't matter
	// because we are ready to read the value from the channel
	go func() {
		result := <-c
		fmt.Println(result)
	}()

	// this will block the execution, this is why the deadlock! error appears
	c <- 1
}

// buffered channel as traffic lights
// with this pattern we are creating a buffered channel with a size of 2
// and we are sending 10 values to the channel
// but the channel values are received 2 by 2 and only 2 goroutines are running at the same time
// and when a goroutine finish, the channel is ready to receive another value
func BufferedChannel() {
	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- i    // send the value to the channel
		wg.Add(1) // we add 1 to the wait group counter
		go doSome(i, &wg, c)
	}

	// when the channel is full, the goroutines are blocked
	// and the for continues when a space on the channel is available

	wg.Wait()
}

func doSome(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("End %d\n", i)

	// when all the logic is done the channels is released to receive another value
	<-c
}
