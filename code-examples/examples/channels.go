package examples

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine
// and receive those values into another goroutine.

func ChannelAnonymousFunc() {
	// Create a new channel with make(chan val-type).
	// Channels are typed by the values they transmit.
	messages := make(chan string)

	// we send a value into a channel using the channel <- syntax.
	go func() {
		// we send "ping" to the messages channel from a new goroutine
		messages <- "ping"
	}()

	// The <-channel syntax receives a value from the channel.
	// we receive the "ping" message we sent above and print it out
	result := <-messages
	fmt.Println(result)

	// By default sends and receives block until both the sender
	// and receiver are ready. This property allowed us to wait at the
	// end of our program for the "ping" message without having to use any
	// other synchronization.
}

// a good practice is to define an entry or exit channel
// we are going to send a message to the channel
// entry channel
func chanSay(text string, c chan<- string) {
	// we send the text to the channel
	c <- text
}

// exit channel
// here we are sending the text on the channel to the text variable
// func chanSay(text string, c <-chan string) {
// 	text = <-c
// }

func ChannelDeclaredFunc() {
	// is a good practice to define the channel size
	c := make(chan string, 1)
	fmt.Println("Sending message")

	// we need to wait the execution of this go routine
	go chanSay("END", c)

	// we do this extracting the value from the channel
	fmt.Println(<-c)
}

func message(text string, c chan<- string) {
	c <- text
}

func ChannelsSelectClose() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"

	// the len function return the number of goroutines on the channel
	// the cap function return the capacity of the channel
	fmt.Println(len(c), cap(c))

	/* close
	close tell to the go runtime that the channel will be closed
	and no more values will be sent to the channel
	ideally the channels must be closed after being used */
	close(c)

	// we can iterate over the channel
	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("message 1", email1)
	go message("message 2", email2)

	// we use the select to receive the messages from the channels
	// the select is like a switch but for channels
	// the select will receive the first message that is available
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email received from email1", m1)
		case m2 := <-email2:
			fmt.Println("Email received from email2", m2)
		}
	}
}
