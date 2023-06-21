// https://leetcode.com/problems/k-radius-subarray-averages/
package random

import "fmt"

func getAverages(nums []int, k int) []int {
	if k == 0 {
		return nums
	}
	
	size := len(nums)
	windowSize := (k * 2) + 1

	avgs := make([]int, size)
	for i := 0; i < size; i++ {
		avgs[i] = -1
	}

	if windowSize > size {
		return avgs
	}
	
	windowSum := 0
	for i := 0; i < windowSize; i++ {
		windowSum += nums[i]
	}

	avgs[k] = windowSum / windowSize

	for i := windowSize; i < size; i++ {
		windowSum = windowSum - nums[i - windowSize] + nums[i]
		avgs[i - k] = windowSum / windowSize
	}

	return avgs
}

func DoKRadiusAverages() {
	fmt.Println(getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
	fmt.Println(getAverages([]int{100000}, 0))
}
