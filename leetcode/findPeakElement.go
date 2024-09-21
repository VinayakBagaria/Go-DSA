// https://leetcode.com/problems/find-peak-element/description/
package leetcode

import "fmt"

func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func DoFindPeakElement() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
