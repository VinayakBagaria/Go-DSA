// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/
package leetcode

import (
	"container/heap"
	"fmt"
)

// [sum, firstIndex, secondIndex]
type closestPairs [][3]int

func (h closestPairs) Len() int           { return len(h) }
func (h closestPairs) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h closestPairs) Less(i, j int) bool { return h[i][0] < h[j][0] }

func (h *closestPairs) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *closestPairs) Pop() interface{} {
	popped := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return popped
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	mh := closestPairs{}

	for i, v := range nums1 {
		mh = append(mh, [3]int{v + nums2[0], i, 0})
	}

	heap.Init(&mh)

	pairs := make([][]int, 0, k)

	for mh.Len() > 0 && k > 0 {
		top := heap.Pop(&mh).([3]int)
		firstVal := nums1[top[1]]
		secondIndex := top[2]
		pairs = append(pairs, []int{firstVal, nums2[secondIndex]})

		if secondIndex < len(nums2)-1 {
			secondIndex++
			heap.Push(&mh, [3]int{firstVal + nums2[secondIndex], top[1], secondIndex})
		}

		k--
	}

	return pairs
}

func DoKPairsClosestSum() {
	fmt.Println("kPairsClosestSum")
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
}
