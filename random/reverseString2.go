// https://leetcode.com/problems/reverse-string-ii/description/
package random

import "fmt"

func reverse(str string) string {
	result := ""
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func reverseStr(s string, k int) string {
	newString := ""
	stringLen := len(s)

	for i := 0; i < stringLen; i = i + (2 * k) {
		if i+k <= stringLen-1 {
			newString += reverse(s[i : i+k])
			remaining := i + (2 * k)
			if remaining <= stringLen-1 {
				newString += s[i+k : remaining]
			} else {
				newString += s[i+k:]
			}
		} else {
			newString += reverse(s[i:])
		}
	}

	return newString
}

func DoReverseString2() {
	fmt.Println(reverseStr("abcdefg", 2))
}
