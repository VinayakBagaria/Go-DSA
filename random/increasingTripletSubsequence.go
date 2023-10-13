// https://leetcode.com/problems/increasing-triplet-subsequence/description/
package random

import "fmt"

func increasingTriplet(nums []int) []int {
	size := len(nums)
	if size < 3 {
		return []int{}
	}

	lMin := make([]int, size)
	lMin[0] = nums[0]
	for i := 1; i < size; i++ {
		lMin[i] = min(lMin[i-1], nums[i])
	}

	rMax := make([]int, size)
	rMax[size-1] = nums[size-1]
	for i := size - 2; i >= 0; i-- {
		rMax[i] = max(rMax[i+1], nums[i])
	}

	for i := 1; i < size-1; i++ {
		if lMin[i-1] < nums[i] && nums[i] < rMax[i+1] {
			return []int{lMin[i-1], nums[i], rMax[i+1]}
		}
	}

	return []int{}
}

func DoIncreasingTripletSubsequence() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}
