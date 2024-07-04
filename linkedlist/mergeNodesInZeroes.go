// https://leetcode.com/problems/merge-nodes-in-between-zeros
package linkedlist

func mergeNodes(head *Node[int]) *Node[int] {
	head = head.Next
	if head == nil {
		return nil
	}

	temp := head
	sum := 0
	for temp.Val != 0 {
		sum += temp.Val
		temp = temp.Next
	}

	head.Val = sum
	head.Next = mergeNodes(temp)
	return head
}

func DoMergeNodesInZeroes() {
	ll := NewLinkedlist[int]()

	nums := []int{0, 3, 1, 0, 4, 5, 2, 0}
	for _, num := range nums {
		ll.addToLast(num)
	}

	mergedLL := NewLinkedlist[int]()
	mergedLL.head = mergeNodes(ll.head)
	mergedLL.print()
}
