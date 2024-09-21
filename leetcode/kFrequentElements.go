// https://leetcode.com/problems/top-k-frequent-elements/description/
package leetcode

import (
	"container/heap"
)

type numInfo struct {
	num  int
	freq int
}

type elementMaxHeap []numInfo

func (h elementMaxHeap) Len() int { return len(h) }

func (h elementMaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h elementMaxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }

func (h *elementMaxHeap) Push(x interface{}) { *h = append(*h, x.(numInfo)) }

func (h *elementMaxHeap) Pop() interface{} {
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

	mh := elementMaxHeap([]numInfo{})
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
