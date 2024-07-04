package linkedlist

import (
	"fmt"
)

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedlist[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{Val: val}
}

func (l *LinkedList[T]) addToFirst(val T) {
	newNode := NewNode(val)
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.Next = l.head
		l.head = newNode
	}
}

func (l *LinkedList[T]) addToLast(val T) {
	newNode := NewNode(val)
	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
}

func (l *LinkedList[T]) print() {
	if l.head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%v -> ", currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Printf("null\n")
}

func (l *LinkedList[T]) deleteFirst() {
	if l.head == nil {
		return
	}

	l.head = l.head.Next
}

func (l *LinkedList[T]) deleteLast() {
	if l.head == nil {
		return
	}

	if l.head.Next == nil {
		l.head = nil
		return
	}

	secondLast := l.head
	lastNode := l.head.Next

	for lastNode.Next != nil {
		secondLast = secondLast.Next
		lastNode = lastNode.Next
	}

	secondLast.Next = nil
}

func (l *LinkedList[T]) reverse() {
	if l.head == nil {
		return
	}

	previous := l.head
	current := l.head.Next

	for current != nil {
		next := current.Next
		current.Next = previous

		previous = current
		current = next
	}
	l.head.Next = nil
	l.head = previous
}

func reverseRecursive[T any](head *Node[T]) *Node[T] {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func DoExample() {
	ll := NewLinkedlist[string]()
	ll.addToFirst("a")
	ll.addToFirst("is")
	ll.addToLast("list")
	ll.addToFirst("this")
	ll.print()
	fmt.Println("delete first ----")
	ll.deleteFirst()
	ll.print()
	fmt.Println("delete last ----")
	ll.deleteLast()
	ll.print()
	fmt.Println("insert again ----")
	ll.addToLast("list")
	ll.addToFirst("this")
	ll.print()
	fmt.Println("Reverse ----")
	ll.reverse()
	ll.print()
	fmt.Println("Reverse Recursive ----")
	ll.head = reverseRecursive(ll.head)
	ll.print()
}
