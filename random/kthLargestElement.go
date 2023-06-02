// https://leetcode.com/problems/kth-largest-element-in-an-array/
package random

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	popped := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return popped
}

func findKthLargest(nums []int, k int) int {
	minHeap := MaxHeap(nums)
	heap.Init(&minHeap)

	popped := 0

	for i := 0; i < k; i++ {
		popped = heap.Pop(&minHeap).(int)
	}

	return popped
}

func DoKthLargestElement() {
	fmt.Println("kth Larget element")
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
