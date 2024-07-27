package leetcode

import (
	"unicode"
)

func decodeString(s string) string {
	var stack []rune
	for _, ch := range s {
		if ch == ']' {
			var decodedString []rune
			// get the encoded string
			for stack[len(stack)-1] != '[' {
				decodedString = append([]rune{stack[len(stack)-1]}, decodedString...)
				stack = stack[:len(stack)-1]
			}
			// pop [ from stack
			stack = stack[:len(stack)-1]

			// get the number k
			base, k := 1, 0
			for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {
				num := int(stack[len(stack)-1] - '0')
				k += num * base
				base *= 10
				stack = stack[:len(stack)-1]
			}

			// decode k[decodedString], by pushing decodedString k times into stack
			for k > 0 {
				stack = append(stack, decodedString...)
				k--
			}
		} else {
			// push the current character to stack
			stack = append(stack, ch)
		}
	}
	// get the result from stack
	return string(stack)
}
