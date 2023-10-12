package main

import (
	"fmt"

	"github.com/j-w-park/go-playground/src/example/fuzzing"
)

func do() {
	// input := "The quick brown fox jumped over the lazy dog"
	// input := "ѹ"
	input := "\xed"
	rev, revErr := fuzzing.Reverse(input)
	doubleRev, doubleRevErr := fuzzing.Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

func main() {
	fmt.Println("main start")
	do()
	fmt.Println("main end")
}
