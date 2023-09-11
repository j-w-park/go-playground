package concurrency

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for i, v := range s {
		time.Sleep(1000 * time.Millisecond)
		sum += v
		fmt.Println(i, v, sum)
	}
	c <- sum // send sum to c
}

func ExampleChannels1() {
	fmt.Println("run ExampleChannels1")
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func ExampleChannels2() {
	fmt.Println("run ExampleChannels2")
	messages := make(chan string)
	go func() {
		time.Sleep(2000 * time.Millisecond)
		messages <- "ping"
	}()
	fmt.Println(<-messages)
}
