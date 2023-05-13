package heap

import "fmt"

func insertElement(heap []int, value int) []int {
	heap = append(heap, value)

	i := len(heap) - 1

	for i > 0 {
		parent := i / 2

		if heap[parent] < heap[i] {
			// swap parent & current index
			temp := heap[parent]
			heap[parent] = heap[i]
			heap[i] = temp

			// next time, we will go up from this node now
			i = parent
		} else {
			break
		}
	}

	return heap
}

func StartInsert() {
	heap := []int{}
	insertions := []int{50, 30, 40, 35}
	for _, node := range insertions {
		heap = insertElement(heap, node)
	}
	fmt.Printf("Max Heap with node order: %v\nPost Insertion becomes Heap: %v\n", insertions, heap)
}
