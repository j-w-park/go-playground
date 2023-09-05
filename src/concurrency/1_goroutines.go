package concurrency

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func ExampleGoroutines1() {
	say("direct")                                     // When we run this program, we see the output of the blocking call first, then the output of the two goroutines.
	go say("world")                                   // starts a new goroutine running
	go func(msg string) { fmt.Println(msg) }("going") // You can also start a goroutine for an anonymous function call. The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.
	say("hello")
}

func ExampleGoroutines2() {
	// Main1 함수는 goroutine함수가 끝나기까지 기다리지 않고 바로 종료된다.
	// 따라서 아래 코드에서는 say() goroutine실행 후 sleep을 하지 않으므로
	// "world"가 출력되지 않는다.
	say("hello")
	go say("world")
}
