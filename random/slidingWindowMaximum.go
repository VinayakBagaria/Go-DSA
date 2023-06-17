// https://leetcode.com/problems/sliding-window-maximum/description/
package random

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	l := 0
	// store indices (decreasing order of elements for these indices)
	queue := []int{}
	result := []int{}

	for r := 0; r < len(nums); r++ {
		// Before we append right window,
		// we need to make sure no smaller values exists in the queue.
		// We insert the current element "r" only if rest of them are bigger.
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[r] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, r)

		// Remove left val from window
		if l > queue[0] {
			queue = queue[1:]
		}

		if r+1 >= k {
			result = append(result, nums[queue[0]])
			l++
		}
	}

	return result
}

func DoSlidingWindowMaximum() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
