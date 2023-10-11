package main

import (
	"fmt"

	equivalentbinarytrees "github.com/j-w-park/go-playground/src/exercise/equivalent_binary_trees"
	"golang.org/x/tour/tree"
)

func do() {
	// t := tree.New(1)
	// ch := make(chan int, 10)
	// go equivalentbinarytrees.WalkRecursive(ch, nil, t)
	// for v := range ch {
	// 	fmt.Println(v)
	// }

	t := tree.New(1)
	ch := make(chan int, 10)
	go equivalentbinarytrees.WalkIterative(ch, t)
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("main start")
	do()
	fmt.Println("main end")
}
