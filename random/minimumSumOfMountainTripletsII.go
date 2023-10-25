// https://leetcode.com/problems/minimum-sum-of-mountain-triplets-ii/description/
package random

import (
	"fmt"
	"math"
)

func minimumSum(nums []int) int {
	n := len(nums)

	// min from left side
	left := make([]int, n)
	left[0] = nums[0]
	for i := 1; i < n; i++ {
		left[i] = min(nums[i], left[i-1])
	}

	// min from right side
	right := make([]int, n)
	right[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = min(nums[i], right[i+1])
	}

	result := math.MaxInt

	for j := 1; j < n-1; j++ {
		if left[j-1] < nums[j] && right[j+1] < nums[j] {
			result = min(result, left[j-1]+nums[j]+right[j+1])
		}
	}

	if result == math.MaxInt {
		return -1
	}
	return result
}

func DoMinimumSumOfMountainTripletsII() {
	fmt.Println(minimumSum([]int{8, 6, 1, 5, 3}))
	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}))
	fmt.Println(minimumSum([]int{6, 5, 4, 3, 4, 5}))
}
