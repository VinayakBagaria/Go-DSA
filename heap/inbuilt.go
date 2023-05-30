package heap

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	n := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return n
}

func StartInbuilt() {
	fmt.Println("inbuilt heap")
	mh := minHeap([]int{1, 5, 8, 4, 3, 7})
	heap.Init(&mh)
	fmt.Println(mh)
	heap.Push(&mh, 2)
	fmt.Println(mh)
	fmt.Println(heap.Pop(&mh))
	fmt.Println(mh)
}
