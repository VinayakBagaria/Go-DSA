// https://leetcode.com/problems/sort-characters-by-frequency/description/
package random

import (
	"container/heap"
	"fmt"
)

type charInfo struct {
	char rune
	freq int
}

type charMaxHeap []charInfo

func (h charMaxHeap) Len() int { return len(h) }

func (h charMaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h charMaxHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}

func (h *charMaxHeap) Push(x interface{}) { *h = append(*h, x.(charInfo)) }

func (h *charMaxHeap) Pop() interface{} {
	popped := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return popped
}

func frequencySort(s string) string {
	frequency := make(map[rune]int)
	for _, word := range s {
		freq, ok := frequency[word]
		if !ok {
			freq = 0
		}
		frequency[word] = freq + 1
	}

	mh := charMaxHeap([]charInfo{})
	heap.Init(&mh)

	for num, freq := range frequency {
		heap.Push(&mh, charInfo{num, freq})
	}

	newString := ""

	for mh.Len() > 0 {
		info := heap.Pop(&mh).(charInfo)
		for i := 0; i < info.freq; i++ {
			newString += string(info.char)
		}
	}

	return newString
}

func DoSortCharsByFrequency() {
	fmt.Println("sortCharsByFrequency")
	fmt.Println(frequencySort("tree"))
}
