package concurrency

import "fmt"

func ExampleRangeAndClose1() {
	fmt.Println("ExampleRangeAndClose1")

	fibonacci := func(n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}

	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}

func ExampleRangeAndClose2() {
	fmt.Println("ExampleRangeAndClose2")

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This range iterates over each element as it’s received from queue.
	// Because we closed the channel above, the iteration terminates after
	// receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
