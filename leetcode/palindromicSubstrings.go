// https://leetcode.com/problems/palindromic-substrings/description/
package leetcode

import "fmt"

func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	count := 0

	for gap := 0; gap < n; gap++ {
		for i, j := 0, gap; j < n; i, j = i+1, j+1 {
			isEqual := s[i] == s[j]
			if gap == 0 {
				dp[i][j] = true
			} else if gap == 1 {
				dp[i][j] = isEqual
			} else {
				dp[i][j] = isEqual && dp[i+1][j-1]
			}

			if dp[i][j] {
				count++
			}
		}
	}

	return count
}

func DoPalindromicSubstrings() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}
