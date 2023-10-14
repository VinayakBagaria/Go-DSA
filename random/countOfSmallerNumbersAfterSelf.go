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

func countSmallerBinarySearch(nums []int) []int {
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

type Node struct {
	Left       *Node
	Right      *Node
	EqualCount int
	LessCount  int
	Val        int
}

func NewNode(val int) *Node {
	return &Node{Val: val, EqualCount: 1, LessCount: 0}
}

func (t *Node) insert(val int) int {
	smallerCounts := 0
	root := t

	for root != nil {
		if val < root.Val {
			root.LessCount++
			if root.Left == nil {
				root.Left = NewNode(val)
				break
			} else {
				root = root.Left
			}
		} else if val > root.Val {
			smallerCounts += root.EqualCount + root.LessCount
			if root.Right == nil {
				root.Right = NewNode(val)
				break
			} else {
				root = root.Right
			}
		} else {
			root.EqualCount++
			smallerCounts += root.LessCount
			break
		}
	}

	return smallerCounts
}

func countSmallerTree(nums []int) []int {
	size := len(nums)
	result := make([]int, size)

	node := NewNode(nums[size-1])
	result[size-1] = 0

	for i := size - 2; i >= 0; i-- {
		result[i] = node.insert(nums[i])
	}

	return result
}

func DoCountOfSmallerNumbersAfterSelf() {
	fmt.Println(countSmallerTree([]int{5, 2, 6, 1}))
	fmt.Println(countSmallerTree([]int{-1, -1}))
	fmt.Println(countSmallerTree([]int{4, 6, 5, 2, 5, 1}))
}
