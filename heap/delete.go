package heap

import (
	"fmt"
)

func findMax(heap []int, i int) int {
	largest := i

	l := 2*i + 1
	r := 2*i + 2

	if l < len(heap) && heap[l] > heap[largest] {
		largest = l
	}

	if r < len(heap) && heap[r] > heap[largest] {
		largest = r
	}

	return largest
}

func deleteTop(heap []int) []int {
	lastNodeIndex := len(heap) - 1
	lastNode := heap[lastNodeIndex]
	heap[0] = lastNode
	heap = heap[:lastNodeIndex]

	i := 0

	for i < len(heap) {
		largest := findMax(heap, i)

		if i != largest {
			// swap
			heap[largest], heap[i] = heap[i], heap[largest]

			i = largest
		} else {
			break
		}
	}

	return heap
}

func StartDelete() {
	heap := []int{40, 30, 10, 20, 15}

	heapCopied := append([]int{}, heap...)

	fmt.Printf("Max Heap: %v\nPost Deletion becomes: %v\n", heap, deleteTop(heapCopied))
}
