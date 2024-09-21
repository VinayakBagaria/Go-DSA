// https://leetcode.com/problems/word-break
package leetcode

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

func wordBreakGraph(s string, wordDict []string) bool {
	n := len(s)
	queue := []int{0}
	seen := make(map[int]struct{})
	words := make(map[string]struct{})

	for _, w := range wordDict {
		words[w] = struct{}{}
	}

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]

		if start == n {
			return true
		}

		for end := start + 1; end <= n; end++ {
			if _, ok := seen[end]; ok {
				continue
			}

			cur := s[start:end]
			if _, ok := words[cur]; ok {
				queue = append(queue, end)
				seen[end] = struct{}{}
			}
		}
	}

	return false
}
