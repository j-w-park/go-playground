package array_and_slice

import "fmt"

func Example3LenAndCap() {
	arr0 := []int{}
	arr0 = append(arr0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(len(arr0), cap(arr0)) // 11 10

	arr1 := make([]int, 10)
	arr1 = append(arr1, 0)
	fmt.Println(len(arr1), cap(arr1)) // 11 20

	arr2 := make([]int, 10, 12)
	arr2 = append(arr2, 0)
	fmt.Println(len(arr2), cap(arr2)) // 11 12
}