package array_and_slice

import "fmt"

func Example3LenAndCap() {
	s0 := []int{}
	s0 = append(s0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(len(s0), cap(s0)) // 11 10

	s1 := make([]int, 10)
	s1 = append(s1, 0)
	fmt.Println(len(s1), cap(s1)) // 11 20

	s2 := make([]int, 10, 12)
	s2 = append(s2, 0)
	fmt.Println(len(s2), cap(s2)) // 11 12
}
