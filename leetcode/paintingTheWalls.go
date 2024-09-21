// https://leetcode.com/problems/painting-the-walls/description/
package leetcode

import (
	"fmt"
	"math"
)

func paintWalls(cost []int, time []int) int {
	size := len(cost)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size+1)
	}

	var solve func(int, int) int
	solve = func(i, remain int) int {
		if remain <= 0 {
			return 0
		}

		if i >= size {
			return math.MaxInt
		}

		if dp[i][remain] != 0 {
			return dp[i][remain]
		}

		paintI := solve(i+1, remain-1-time[i])
		if paintI != math.MaxInt {
			paintI += cost[i]
		}
		notPaintI := solve(i+1, remain)
		dp[i][remain] = min(paintI, notPaintI)
		return dp[i][remain]
	}

	return solve(0, size)
}

func DoPaintingTheWalls() {
	fmt.Println(paintWalls([]int{1, 2, 3, 2}, []int{1, 2, 3, 2}))
}
