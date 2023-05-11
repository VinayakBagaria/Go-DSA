// https://leetcode.com/problems/longest-increasing-subsequence/
package dp

import (
	"fmt"
	"go-dsa/utils"
)

func lengthOfLIS(nums []int) int {
	// https://www.youtube.com/watch?v=odrfUCS9sQk
	dp := make([]int, len(nums))
	dp[0] = 1

	for i, num := range nums {
		maximum := 0
		for j := 0; j < i; j++ {
			if nums[j] < num {
				maximum = utils.Max(maximum, dp[j])
			}
		}

		dp[i] = maximum + 1
	}

	return utils.MaxOfArray(dp)
}

func DoLisLength() {
	inputs := [][]int{
		{1, 2, 4, 3},
		{10, 9, 2, 5, 3, 7, 101, 18},
	}

	for _, input := range inputs {
		fmt.Printf("Input: %v\nOutput: %d\n", input, lengthOfLIS(input))
	}
}
