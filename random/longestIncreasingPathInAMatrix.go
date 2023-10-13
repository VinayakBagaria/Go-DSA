// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description/
package random

import (
	"fmt"
	"math"
)

var DIRECTIONS = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	var dfs func(int, int) int
	dfs = func(r, c int) int {
		if dp[r][c] != 0 {
			return dp[r][c]
		}

		maximum := 1
		for _, dir := range DIRECTIONS {
			x := r + dir[0]
			y := c + dir[1]

			if x < 0 || y < 0 || x >= m || y >= n {
				continue
			}

			if matrix[r][c] <= matrix[x][y] {
				continue
			}

			maximum = max(maximum, 1+dfs(x, y))
		}

		dp[r][c] = maximum
		return dp[r][c]
	}

	result := math.MinInt

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			result = max(result, dfs(r, c))
		}
	}

	return result
}

func DoLongestIncreasingPathInAMatrix() {
	fmt.Println(longestIncreasingPath([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}))
	fmt.Println(longestIncreasingPath([][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}))
}
