// https://leetcode.com/problems/k-closest-points-to-origin/description/
package random

import (
	"container/heap"
)

type pointInfo struct {
	point    []int
	distance int
}

type minHeap []pointInfo

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pointInfo))
}

func (h *minHeap) Pop() interface{} {
	popped := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return popped
}

func kClosest(points [][]int, k int) [][]int {
	mh := minHeap([]pointInfo{})
	heap.Init(&mh)

	for _, point := range points {
		distance := (point[0] * point[0]) + (point[1] * point[1])
		heap.Push(&mh, pointInfo{
			point,
			distance,
		})
	}

	answers := [][]int{}

	for i := 0; i < k; i++ {
		information := heap.Pop(&mh).(pointInfo)
		answers = append(answers, information.point)
	}

	return answers
}
