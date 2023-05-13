package heap

import (
	"fmt"
)

type MaxHeap struct {
	heap []int
}

func (h *MaxHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *MaxHeap) heapify(i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	size := len(h.heap)

	if l < size && h.heap[l] > h.heap[largest] {
		largest = l
	}
	if r < size && h.heap[r] > h.heap[largest] {
		largest = r
	}

	if i != largest {
		h.swap(i, largest)
		h.heapify(largest)
	}
}

func StartHeapify() {
	arr := []int{20, 10, 30, 5, 50, 40}
	h := &MaxHeap{heap: append([]int{}, arr...)}

	for i := len(h.heap) / 2; i >= 0; i-- {
		h.heapify(i)
	}
	fmt.Printf("Initial array: %v\nPost Heapify, Max Heap becomes: %v\n", arr, h.heap)
}
