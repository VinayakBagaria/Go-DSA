// https://leetcode.com/problems/majority-element/description/
// Boyer-Moore algorithm
package leetcode

import "fmt"

func majorityElement(nums []int) int {
	candidate := -1
	count := 0

	for _, n := range nums {
		if count == 0 {
			candidate = n
			count++
		} else if candidate == n {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func DoMajorityElement() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
