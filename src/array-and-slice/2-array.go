package array_and_slice

import "fmt"

func Example2Array() {
	var arr0 = [3]int{1, 2, 3}
	var arr1 = [...]int{1, 2, 3}
	fmt.Println(arr0 == arr1) // true

	var arr2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2 == arr3) // true
}
