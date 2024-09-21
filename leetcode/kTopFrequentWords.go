// https://leetcode.com/problems/top-k-frequent-words/description/
package leetcode

import (
	"container/heap"
	"fmt"
)

type wordInfo struct {
	word string
	freq int
}

type wordMaxHeap []wordInfo

func (h wordMaxHeap) Len() int { return len(h) }

func (h wordMaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h wordMaxHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].word < h[j].word
	}
	return h[i].freq > h[j].freq
}

func (h *wordMaxHeap) Push(x interface{}) { *h = append(*h, x.(wordInfo)) }

func (h *wordMaxHeap) Pop() interface{} {
	popped := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return popped
}

func topKFrequentWords(words []string, k int) []string {
	frequencies := make(map[string]int)
	for _, word := range words {
		freq, ok := frequencies[word]
		if !ok {
			freq = 0
		}
		frequencies[word] = freq + 1
	}

	mh := wordMaxHeap([]wordInfo{})
	heap.Init(&mh)

	for num, freq := range frequencies {
		heap.Push(&mh, wordInfo{num, freq})
	}

	elements := []string{}

	for i := 0; i < k; i++ {
		info := heap.Pop(&mh).(wordInfo)
		elements = append(elements, info.word)
	}

	return elements
}

func DoTopKFrequentWords() {
	fmt.Println("topKFrequentWords")
	fmt.Println(topKFrequentWords([]string{"i", "love", "leetcode", "i", "love", "coding"}, 1))
	fmt.Println(topKFrequentWords([]string{"i", "love", "leetcode", "i", "love", "coding"}, 3))
	fmt.Println(topKFrequentWords([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
}
