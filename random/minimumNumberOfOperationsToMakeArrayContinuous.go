// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/description/
package random

import (
	"fmt"
	"math"
	"sort"
)

func minOperations(nums []int) int {
	actualLength := len(nums)
	sort.Ints(nums)

	set := []int{nums[0]}
	for i := 1; i < actualLength; i++ {
		if nums[i] != nums[i-1] {
			set = append(set, nums[i])
		}
	}
	uniqueLength := len(set)

	result := math.MaxInt
	r := 0
	for l := range set {
		windowEnd := set[l] + actualLength - 1
		for r < uniqueLength && set[r] <= windowEnd {
			r++
		}

		windowLength := r - l
		result = min(result, actualLength-windowLength)
	}

	return result
}

func DoMinimumNumberOfOperationsToMakeArrayContinuous() {
	fmt.Println(minOperations([]int{4, 2, 5, 3}))
	fmt.Println(minOperations([]int{1, 2, 3, 5, 6}))
}
