package fuzzing

import (
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", fmt.Errorf("invalid utf-8 string: %q", s)
	}
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-(i+1)] = r[len(r)-(i+1)], r[i]
	}
	return string(r), nil
}

// failed with: "\xed"
// func Reverse(s string) string {
// 	r := []rune(s)
// 	for i := 0; i < len(r)/2; i++ {
// 		r[i], r[len(r)-(i+1)] = r[len(r)-(i+1)], r[i]
// 	}
// 	return string(r)
// }

// failed with: "00000000ѹ"
// func Reverse(s string) string {
// 	b := []byte(s)
// 	for i := 0; i < len(b)/2; i++ {
// 		b[i], b[len(b)-(i+1)] = b[len(b)-(i+1)], b[i]
// 	}
// 	return string(b)
// }
