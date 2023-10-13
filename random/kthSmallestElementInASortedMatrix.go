// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/submissions/
package random

import (
	"container/heap"
	"fmt"
)

func kthSmallestElement(matrix [][]int, k int) int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for _, row := range matrix {
		for _, curr := range row {
			heap.Push(maxHeap, curr)
			if maxHeap.Len() > k {
				heap.Pop(maxHeap)
			}
		}
	}

	return (*maxHeap)[0]
}

func DoKthSmallestElementInASortedMatrix() {
	fmt.Println(kthSmallestElement([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}, 8))
}
