// https://leetcode.com/problems/top-k-frequent-elements/description/
package random

import (
	"container/heap"
)

type numInfo struct {
	num  int
	freq int
}

type maxHeap []numInfo

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Less(i, j int) bool  { return h[i].freq > h[j].freq }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(numInfo)) }
func (h *maxHeap) Pop() interface{} {
	popped := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return popped
}

func topKFrequent(nums []int, k int) []int {
	frequencies := make(map[int]int)
	for _, num := range nums {
		freq, ok := frequencies[num]
		if !ok {
			freq = 0
		}
		frequencies[num] = freq + 1
	}

	mh := maxHeap([]numInfo{})
	heap.Init(&mh)

	for num, freq := range frequencies {
		heap.Push(&mh, numInfo{num, freq})
	}

	elements := []int{}

	for i := 0; i < k; i++ {
		info := heap.Pop(&mh).(numInfo)
		elements = append(elements, info.num)
	}

	return elements
}
