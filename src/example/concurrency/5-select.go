package concurrency

import (
	"fmt"
	"time"
)

func ExampleSelect1() {
	fibonacci := func(c, quit chan int) {
		n, x, y := 0, 0, 1
		for {
			// The select statement lets a goroutine wait on multiple
			// communication operations.
			// A select blocks until one of its cases can run, then it executes
			// that case. It chooses one at random if multiple are ready.
			select {
			case c <- x:
				fmt.Printf("send (%d, %d)\n", n, x)
				x, y = y, x+y
			case <-quit:
				fmt.Printf("quit (%d, %d, %d)\n", n, x, y)
				return
			// The default case in a select is run if no other case is ready.
			// Use a default case to try a send or receive without blocking:
			default:
				n++
			}
		}
	}

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("received (%d, %d)\n", i, <-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

func ExampleSelect2() {
	tick := time.NewTicker(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick.C:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
