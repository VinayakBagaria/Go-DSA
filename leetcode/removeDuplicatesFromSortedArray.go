// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
package leetcode

import "fmt"

func removeDuplicates(nums []int) int {
	j := 0

	for i, current := range nums {
		if i > 0 && current == nums[i-1] {
			continue
		} else {
			nums[j] = current
			j++
		}
	}

	return j
}

func DoRemoveDuplicatesFromSortedArray() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
