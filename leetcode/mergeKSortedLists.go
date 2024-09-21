// https://leetcode.com/problems/merge-k-sorted-lists/description/
package leetcode

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type nodeHeap []*ListNode

func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h *nodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *nodeHeap) Pop() interface{} {
	popped := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return popped
}

func mergeKLists(lists []*ListNode) *ListNode {
	temp := &ListNode{}
	head := temp

	mh := nodeHeap([]*ListNode{})
	heap.Init(&mh)

	for _, listNode := range lists {
		if listNode != nil {
			heap.Push(&mh, listNode)
		}
	}

	for mh.Len() > 0 {
		node := heap.Pop(&mh).(*ListNode)
		temp.Next = node
		temp = temp.Next
		if node.Next != nil {
			heap.Push(&mh, node.Next)
		}
	}

	return head.Next
}

func DoMergeKSortedLists() {
	fmt.Println("DoMergeKSortedLists")

	lists := []*ListNode{}

	elements := [][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}
	for _, eachElementList := range elements {
		temp := &ListNode{Val: eachElementList[0], Next: nil}
		head := temp

		for i := 1; i < len(eachElementList); i++ {
			temp.Next = &ListNode{Val: eachElementList[i], Next: nil}
			temp = temp.Next
		}

		lists = append(lists, head)
	}

	printList(mergeKLists(lists))
}

func printList(node *ListNode) {
	head := node

	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}

	fmt.Println("nil")
}
