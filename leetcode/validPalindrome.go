// https://leetcode.com/problems/valid-palindrome/
package leetcode

import (
	"fmt"
	"strings"
	"unicode"
)

func checkValid(cb byte) bool {
	c := rune(cb)
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func isPalindrome(s string) bool {
	l := len(s)
	start := 0
	last := l - 1

	for start <= last {
		for start < last && !checkValid(s[start]) {
			start++
		}
		for last > start && !checkValid(s[last]) {
			last--
		}
		if !strings.EqualFold(string(s[start]), string(s[last])) {
			return false
		}
		start++
		last--
	}

	return true
}

func DoValidPalindrome() {
	fmt.Println("validPalindrome")
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("race a car"))
}
