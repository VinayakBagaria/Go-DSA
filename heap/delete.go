package heap

import (
	"fmt"
)

func findMax(heap []int, i int) int {
	largestIndex := i

	leftIndex := 2 * i
	rightIndex := 2*i + 1

	if leftIndex < len(heap) && heap[leftIndex] > heap[largestIndex] {
		largestIndex = leftIndex
	}

	if rightIndex < len(heap) && heap[rightIndex] > heap[largestIndex] {
		largestIndex = rightIndex
	}

	return largestIndex
}

func deleteTop(heap []int) []int {
	lastNodeIndex := len(heap) - 1
	lastNode := heap[lastNodeIndex]
	heap[0] = lastNode
	heap = heap[:lastNodeIndex]

	i := 0

	for i < len(heap) {
		largerIndex := findMax(heap, i)

		if i != largerIndex {
			// swap
			heap[largerIndex], heap[i] = heap[i], heap[largerIndex]

			i = largerIndex
		} else {
			break
		}
	}

	return heap
}

func StartDelete() {
	heap := []int{40, 30, 10, 20, 15}
	var heapCopied []int
	heapCopied = append(heapCopied, heap...)

	fmt.Printf("Max Heap: %v\nPost Deletion becomes: %v\n", heap, deleteTop(heapCopied))
}
