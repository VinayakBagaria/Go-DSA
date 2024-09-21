// https://leetcode.com/problems/longest-palindromic-substring/description/
package leetcode

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	result := ""

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

			if dp[i][j] && len(result) < j-i+1 {
				result = s[i : j+1]
			}
		}
	}

	return result
}

func DoLongestPalindromicSubstring() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("aaa"))
}
