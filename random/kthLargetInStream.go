// https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
package random

import (
	"container/heap"
	"fmt"
)

type heaper []int

func (h heaper) Less(i, j int) bool { return h[i] < h[j] }

func (h heaper) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h heaper) Len() int { return len(h) }

func (h *heaper) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *heaper) Pop() interface{} {
	popped := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return popped
}

type KthLargest struct {
	k       int
	minHeap heaper
}

func constructor(k int, nums []int) KthLargest {
	minHeap := heaper(nums)
	heap.Init(&minHeap)
	return KthLargest{k: k, minHeap: minHeap}
}

func (this *KthLargest) add(val int) int {
	heap.Push(&this.minHeap, val)
	for this.minHeap.Len() > this.k {
		heap.Pop(&this.minHeap)
	}
	return this.minHeap[0]
}

func DoKthLargestInStream() {
	fmt.Println("kthLargestInStream")
	obj := constructor(3, []int{4, 5, 8, 2})
	fmt.Println(obj.add(3))
	fmt.Println(obj.add(5))
	fmt.Println(obj.add(10))
	fmt.Println(obj.add(9))
}
