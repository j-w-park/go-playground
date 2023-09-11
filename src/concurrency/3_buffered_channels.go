package concurrency

import "fmt"

func ExampleBufferedChannels1() {
	fmt.Println("run ExampleBufferedChannels1")
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	// messages <- "channel" // will cause deadlock
	fmt.Println(<-messages, <-messages)
}
