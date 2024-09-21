// https://leetcode.com/problems/longest-common-subsequence/description/
package leetcode

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i >= n || j >= m {
			return 0
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		ans := 0
		if text1[i] == text2[j] {
			ans = 1 + dfs(i+1, j+1)
		} else {
			ans = max(dfs(i, j+1), dfs(i+1, j))
		}

		dp[i][j] = ans
		return dp[i][j]
	}

	return dfs(0, 0)
}

func DoLongestCommonSubsequence() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
}
