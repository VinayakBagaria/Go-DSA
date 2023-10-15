// https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/description/
package random

import "fmt"

var MOD = int(1e9 + 7)

func numWays(steps int, arrLen int) int {
	arrLen = min(steps, arrLen)

	dp := make([][]int, arrLen)
	for i := range dp {
		dp[i] = make([]int, steps+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(index, remaining int) int {
		if remaining == 0 {
			if index == 0 {
				return 1
			}
			return 0
		}

		if index < 0 || index >= arrLen {
			return 0
		}

		if dp[index][remaining] != -1 {
			return dp[index][remaining]
		}

		right := index + 1
		left := index - 1
		stay := index
		branches := []int{right, left, stay}

		ways := 0
		for _, branch := range branches {
			ways += dfs(branch, remaining-1)
			ways %= MOD
		}
		dp[index][remaining] = ways
		return ways
	}

	return dfs(0, steps)
}

func DoNumberOfWaysToStayInTheSamePlaceAfterSomeSteps() {
	fmt.Println(numWays(3, 2))
	fmt.Println(numWays(2, 4))
	fmt.Println(numWays(4, 2))
}
