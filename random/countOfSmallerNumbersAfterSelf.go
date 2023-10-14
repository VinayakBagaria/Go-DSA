// https://leetcode.com/problems/count-of-smaller-numbers-after-self/description/
package random

import "fmt"

func bisectLeft(sorted []int, target int) int {
	l := 0
	r := len(sorted)

	for l < r {
		mid := l + (r-l)/2

		if sorted[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func addToSorted(sorted []int, index, target int) []int {
	// Expand the slice to accommodate the new element
	sorted = append(sorted, 0)
	// Shift elements to the right to make room for the new number
	copy(sorted[index+1:], sorted[index:])
	sorted[index] = target
	return sorted
}

func countSmaller(nums []int) []int {
	size := len(nums)
	result := make([]int, size)
	sorted := []int{}

	for i := size - 1; i >= 0; i-- {
		curr := nums[i]
		index := bisectLeft(sorted, curr)
		sorted = addToSorted(sorted, index, curr)
		result[i] = index
	}

	return result
}

func DoCountOfSmallerNumbersAfterSelf() {
	// fmt.Println(bisectLeft([]int{1, 6}, 2))
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
}
