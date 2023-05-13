package sorting

type MaxHeap struct {
	heap []int
}

func (h *MaxHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *MaxHeap) heapify(i, size int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < size && h.heap[l] > h.heap[largest] {
		largest = l
	}

	if r < size && h.heap[r] > h.heap[largest] {
		largest = r
	}

	if i != largest {
		h.swap(i, largest)
		h.heapify(largest, size)
	}
}

func HeapSort(items []int) []int {
	h := &MaxHeap{heap: append([]int{}, items...)}

	size := len(h.heap)

	for i := size / 2; i >= 0; i-- {
		h.heapify(i, size)
	}

	for i := size - 1; i >= 0; i-- {
		h.swap(0, i)
		h.heapify(0, i)
	}

	return h.heap
}
