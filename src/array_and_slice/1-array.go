package example_array_and_slice

import "fmt"

func Example1Array() {
	var a0 = [3]int{1, 2, 3}
	var a1 = [...]int{1, 2, 3}
	fmt.Println(a0 == a1) // true

	var a2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var a3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(a2 == a3) // true
}
