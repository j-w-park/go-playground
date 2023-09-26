package type_and_interface

import (
	"fmt"
)

func TypeAssertion() {
	type Number int

	var (
		v any    = 1
		n Number = 20
	)
	v = n
	// fmt.Println(v.(int) + 1) // panic: v is Number, not int
	fmt.Println(v.(Number) + 1)
}
