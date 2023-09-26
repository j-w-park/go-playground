package array_and_slice

import "fmt"

func Example2Slice() {
	{
		var s0 = []int{1, 2, 3}
		var s1 = []int{}
		var s2 []int
		fmt.Println("s0:", s0)               // [1 2 3]
		fmt.Println("s1 == nil:", s1 == nil) // false
		fmt.Println("s2 == nil:", s2 == nil) // true
		// fmt.Println(s0 == s1) // invalid operation: slice can only be compared to nil

		var s3 = []int{1, 5: 3, 10, 10: 100}
		fmt.Println("s2:", s3) // [1 0 0 0 0 3 0 0 0 0 100]

		var s4 = [][]int{{1, 2, 3}, {4, 6}}
		fmt.Println("s3:", s4) // [[1 2 3] [4 6]]

		var s5 = append(
			append(
				append(
					append(
						append(
							append(
								make([]int, 0),
								1),
							2, 3),
						4, 5, 6),
					7, 8, 9, 10),
				11, 12, 13, 14, 15),
			16, 17, 18, 19, 20, 21,
		)
		fmt.Println("s5:", s5) // [1 2 3 ... 21]
	}

	{
		fmt.Println("slicing side effect")
		x0 := []int{1, 2, 3, 4}
		y0 := x0[1:3]
		fmt.Println("y0:", y0, len(y0), cap(y0)) // [2 3] 2 3
		y0[0] = 100                              // 메모리를 공유하기 때문에 x0의 값도 바뀐다.
		fmt.Println("x0:", x0)                   // [1 100 3 4]
		fmt.Println("y0:", y0)                   // [100 3]
		y0 = append(y0, 10)                      // y0의 capacity가 3이기 때문에 x0가 가리키는 메모리에 할당된다.
		fmt.Println("x0:", x0)                   // [1 100 3 10]
		fmt.Println("y0:", y0)                   // [100 3 10]

		fmt.Println("full slice expression")
		x1 := []int{1, 2, 3, 4}
		y1 := x1[1:3:3]                          // (start) : (end{exclusive}) : (max-capacity index)
		fmt.Println("y1:", y1, len(y1), cap(y1)) // [2 3] 2 2(= 3 - 1)
		y1[0] = 100                              // 여전히 메모리를 공유하기 때문에 x1의 값도 바뀐다.
		fmt.Println("x1:", x1)                   // [1 100 3 4]
		fmt.Println("y1:", y1)                   // [100 3]
		y1 = append(y1, 10)                      // y1의 capacity가 2이기 때문에 append시 새로운 메모리에 할당되고 x1과는 더이상 메모리를 공유하지 않는다.
		fmt.Println("x1:", x1)                   // [1 100 3 4]
		fmt.Println("y1:", y1)                   // [100 3 10]

		fmt.Println("copy")
		x2 := []int{1, 2, 3, 4}
		y2 := make([]int, 2)
		copy(y2, x2[1:3])
		fmt.Println("y2:", y2, len(y2), cap(y2)) // [2 3] 2 2
		y2[0] = 100                              // 메모리를 공유하지 않음
		fmt.Println("x2:", x2)                   // [1 2 3 4]
		fmt.Println("y2:", y2)                   // [100 3]
	}
}
