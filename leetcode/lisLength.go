// https://leetcode.com/problems/longest-increasing-subsequence/description/
package leetcode

import (
	"fmt"
	"go-dsa/utils"
)

type QueueElement struct {
	sequence []int
	index    int
}

func approach1(nums []int) int {
	queue := []QueueElement{}

	longest := -1

	for index, number := range nums {
		sequence := []int{number}
		queue = append(queue, QueueElement{sequence: sequence, index: index})
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		longest = utils.Max(longest, len(node.sequence))

		lastElement := node.sequence[len(node.sequence)-1]
		index := node.index

		for i := index; i < len(nums); i++ {
			if nums[i] > lastElement {
				newSequence := utils.CopyFromOneArrayToAnother(node.sequence)
				newSequence = append(newSequence, nums[i])
				queue = append(queue, QueueElement{sequence: newSequence, index: i})
			}
		}
	}

	return longest
}

func DoLisLength() {
	inputs := [][]int{
		{1, 2, 4, 3},
		{10, 9, 2, 5, 3, 7, 101, 18},
	}

	for _, input := range inputs {
		fmt.Printf("Input: %v\nOutput: %d\n", input, approach1(input))
	}
}
