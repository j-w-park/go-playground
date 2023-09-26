package type_and_interface

import (
	"fmt"
	"io"
)

func TypeSwitch() {
	doThings := (func(i any) {
		switch j := i.(type) {
		case nil:
			fmt.Println("type: nil")
		case int:
			fmt.Println("type: int")
		case io.Reader:
			fmt.Println("type: io.Reader")
		case string:
			fmt.Println("type: string")
		case bool, rune:
			fmt.Println("type: bool or rune(any)")
		default:
			fmt.Printf("unexpected type %T\n", j)
		}
	})

	doThings(123.2)
}
