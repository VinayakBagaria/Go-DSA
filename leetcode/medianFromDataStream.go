// https://leetcode.com/problems/find-median-from-data-stream
// https://www.youtube.com/watch?v=Yv2jzDzYlp8
package leetcode

import "fmt"

type Heap struct {
	values     []int
	comparator func(a, b int) bool
}

func (h *Heap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *Heap) Add(num int) {
	h.values = append(h.values, num)

	i := h.Size() - 1

	for i > 0 {
		parent := (i - 1) / 2

		if h.comparator(h.values[parent], h.values[i]) {
			break
		} else {
			h.Swap(parent, i)
			i = parent
		}
	}
}

func (h *Heap) Poll() int {
	firstNode := h.Peek()
	lastIndex := h.Size() - 1
	h.values[0] = h.values[lastIndex]
	h.values = h.values[:lastIndex]

	i := 0
	for i < h.Size() {
		sizest := i
		l := 2*i + 1
		r := 2*i + 2

		if l < h.Size() && h.comparator(h.values[l], h.values[sizest]) {
			sizest = l
		}
		if r < h.Size() && h.comparator(h.values[r], h.values[sizest]) {
			sizest = r
		}

		if i == sizest {
			break
		} else {
			h.Swap(i, sizest)
			i = sizest
		}
	}

	return firstNode
}

func (h *Heap) Peek() int {
	return h.values[0]
}

func (h *Heap) Size() int {
	return len(h.values)
}

type MedianFinder struct {
	maxHeap *Heap
	minHeap *Heap
}

func Constructor() MedianFinder {
	maxHeap := Heap{values: []int{}, comparator: func(a, b int) bool { return a > b }}
	minHeap := Heap{values: []int{}, comparator: func(a, b int) bool { return a < b }}
	return MedianFinder{maxHeap: &maxHeap, minHeap: &minHeap}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.maxHeap.Size() == 0 || mf.maxHeap.Peek() >= num {
		mf.maxHeap.Add(num)
	} else {
		mf.minHeap.Add(num)
	}

	if mf.maxHeap.Size() > mf.minHeap.Size()+1 {
		mf.minHeap.Add(mf.maxHeap.Poll())
	} else if mf.maxHeap.Size() < mf.minHeap.Size() {
		mf.maxHeap.Add(mf.minHeap.Poll())
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxHeap.Size() == mf.minHeap.Size() {
		return float64(mf.maxHeap.Peek()+mf.minHeap.Peek()) / float64(2)
	}

	return float64(mf.maxHeap.Peek())
}

func DoMedianFromDataStream() {
	obj := Constructor()
	nums := []int{155, 66, 114, 0, 60, 73, 109, 26, 154, 0, 107}
	for _, n := range nums {
		obj.AddNum(n)
	}
	fmt.Printf("Median of array: %v is %f\n", nums, obj.FindMedian())
}
