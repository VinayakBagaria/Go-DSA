// https://leetcode.com/problems/kth-largest-element-in-an-array/
package random

import "fmt"

type MinHeap struct {
	heap []int
}

func (h *MinHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *MinHeap) heapify(i int) {
	smallest := i
	l := 2*i + 1
	r := 2*i + 2
	size := len(h.heap)

	if l < size && h.heap[l] < h.heap[smallest] {
		smallest = l
	}

	if r < size && h.heap[r] < h.heap[smallest] {
		smallest = r
	}

	if i != smallest {
		h.swap(i, smallest)
		h.heapify(smallest)
	}
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{heap: nums[:k]}

	for i := k - 1; i >= 0; i-- {
		h.heapify(i)
	}

	for i := k; i < len(nums); i++ {
		if h.heap[0] < nums[i] {
			h.heap[0] = nums[i]
			h.heapify(0)
		}
	}

	return h.heap[0]
}

func DoKthLargestElement() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	// fmt.Println(findKthLargest([]int{20, 10, 60, 30, 50, 40}, 3))
}
