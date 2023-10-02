// https://leetcode.com/problems/word-break
package random

func wordBreakBottomUp(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)

	dp[n] = true

	for i := n - 1; i >= 0; i-- {
		for _, w := range wordDict {
			wLen := len(w)

			if i+wLen <= n && s[i:i+wLen] == w {
				dp[i] = dp[i+wLen]
			}

			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}

func wordBreakTopDown(s string, wordDict []string) bool {
	n := len(s)
	dp := make(map[int]bool)

	var dfs func(int) bool
	dfs = func(i int) bool {
		if i > n {
			return false
		}

		if i == n {
			return true
		}

		if val, ok := dp[i]; ok {
			return val
		}

		for _, w := range wordDict {
			wLen := len(w)

			if i+wLen <= n && s[i:i+wLen] == w {
				dp[i] = dfs(i + wLen)
				if dp[i] {
					return true
				}
			}
		}

		return false
	}

	return dfs(0)
}
