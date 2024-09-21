// https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/description/
package leetcode

import (
	"container/heap"
	"fmt"
)

type rowHeap []int

func (h rowHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h rowHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h rowHeap) Len() int { return len(h) }

func (h *rowHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *rowHeap) Pop() interface{} {
	popped := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return popped
}

type information struct {
	mat     [][]int
	k       int
	maxHeap rowHeap
	rows    int
	columns int
}

func (i *information) dfsSum(r, c, currentSum int) {
	currentSum += i.mat[r][c]

	if r == i.rows-1 {
		heap.Push(&i.maxHeap, currentSum)
		for i.maxHeap.Len() > i.k {
			heap.Pop(&i.maxHeap)
		}
		return
	}

	for column := 0; column < i.columns; column++ {
		if column != i.columns {
			i.dfsSum(r+1, column, currentSum)
		}
	}
}

func kthSmallest(mat [][]int, k int) int {
	maxHeap := rowHeap([]int{})

	i := &information{mat: mat, k: k, maxHeap: maxHeap, rows: len(mat), columns: len(mat[0])}
	for column := 0; column < i.columns; column++ {
		i.dfsSum(0, column, 0)
	}

	return i.maxHeap[0]
}

func DoKthSmallestMatrixRowSum() {
	fmt.Println("kthSmallestMatrixRowSum")
	s := kthSmallest([][]int{
		{1, 3, 11},
		{2, 4, 6},
	}, 9)
	fmt.Println(s)
}
