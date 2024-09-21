// https://leetcode.com/problems/minimum-window-substring/description/
package leetcode

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	count := make(map[byte]int)
	for i := range t {
		count[t[i]]++
	}
	need := len(count)

	window := make(map[byte]int)
	have := 0

	minResult := math.MaxInt
	minAnswer := ""

	l := 0
	for r := range s {
		rightChar := s[r]

		window[rightChar]++
		if counted, ok := count[rightChar]; ok && window[rightChar] == counted {
			have++
		}

		for have == need {
			windowLength := r - l + 1
			if windowLength < minResult {
				minResult = windowLength
				minAnswer = s[l : r+1]
			}

			leftChar := s[l]
			l++
			window[leftChar]--
			if counted, ok := count[leftChar]; ok && window[leftChar] < counted {
				have--
			}
		}
	}

	return minAnswer
}

func DoMinWindow() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("aa", "aa"))
	fmt.Println(minWindow("a", "aa"))
}
