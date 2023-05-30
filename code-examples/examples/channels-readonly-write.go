package examples

import "fmt"

// when the arrow points to the channel
// c CHAN <- int
// write only  channel

// when the arrow moves away from the channel
// in <- CHAN int
// read only channels

// write only channel
func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

// in -  read only channel
// out - write only channel
func Double(in <-chan int, out chan<- int) {
	// we iterate over the read only channel values
	for value := range in {
		// we send the value multiplied by 2 to the write only channel
		out <- 2 * value
	}

	// and the channel is closed
	close(out)
}

// read only channels
func Print(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func ChannelsReadOnlyWrite() {
	generator := make(chan int)
	doubles := make(chan int)

	// write the numbers from 1 to 10 to the generator channel
	go Generator(generator)

	// send the values from the generator channel to the doubles channel
	go Double(generator, doubles)

	// read all the values on the doubles channel
	Print(doubles)
}
