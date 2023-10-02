// https://leetcode.com/problems/word-break
package random

func wordBreak(s string, wordDict []string) bool {
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
