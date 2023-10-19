// https://leetcode.com/problems/backspace-string-compare/description/
package random

import "fmt"

func makeStack(s string) []byte {
	stack := []byte{}
	for i := range s {
		if s[i] == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return stack
}

func backspaceCompare(s string, t string) bool {
	return string(makeStack(s)) == string(makeStack(t))
}

func DoBackspaceStringCompare() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a#c", "b"))
}
