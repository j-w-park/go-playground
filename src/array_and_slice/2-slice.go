package array_and_slice

import "fmt"

func Example2Slice() {
	var s0 = []int{1, 2, 3}
	fmt.Println(s0) // [1 2 3]

	var s1 = []int{1, 5: 3, 10, 10: 100}
	fmt.Println(s1) // [1 0 0 0 0 3 0 0 0 0 100]
}
