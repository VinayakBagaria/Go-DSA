// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/editorial/
package random

import (
	"fmt"
)

func getMaxUniqueChars(s string) int {
	counter := make(map[rune]struct{})
	for _, c := range s {
		counter[c] = struct{}{}
	}
	return len(counter)
}

func longestSubstringSlidingWindow(s string, k int) int {
	maxUnique := getMaxUniqueChars(s)
	result := 0

	for currUnique := 1; currUnique <= maxUnique; currUnique++ {
		counter := make(map[byte]int)

		windowStart := 0
		windowEnd := 0
		unique := 0
		atleastKCount := 0

		for windowEnd < len(s) {
			ws := s[windowStart]
			we := s[windowEnd]

			if unique <= currUnique {
				counter[we]++
				if counter[we] == 1 {
					unique++
				}
				if counter[we] == k {
					atleastKCount++
				}
				windowEnd++
			} else {
				if counter[ws] == k {
					atleastKCount--
				}
				counter[ws]--
				if counter[ws] == 0 {
					unique--
				}
				windowStart++
			}

			if unique == atleastKCount {
				result = max(result, windowEnd-windowStart)
			}
		}
	}

	return result
}

func helper(s string, k, start, end int) int {
	// store count of every char in the string
	counter := make(map[byte]int)
	for i := start; i < end; i++ {
		counter[s[i]]++
	}

	for mid := start; mid < end; mid++ {
		if counter[s[mid]] >= k {
			continue
		}

		// we got a faulty character now whose count is < k
		// so now search for answer on left side or right side of this char
		return max(helper(s, k, start, mid), helper(s, k, mid+1, end))
	}

	return end - start
}

func longestSubstringDivideAndConquer(s string, k int) int {
	return helper(s, k, 0, len(s))
}

func DoLongestSubstringWithAtleastKRepeatingCharacters() {
	inputs := []struct {
		word string
		k    int
	}{
		{"aaabb", 3},
		{"ababbc", 2},
		{"bbaaacbd", 3},
	}
	for _, input := range inputs {
		fmt.Printf("%s with %d\n", input.word, input.k)
		fmt.Println(longestSubstringDivideAndConquer(input.word, input.k))
		fmt.Println(longestSubstringSlidingWindow(input.word, input.k))
	}
}
