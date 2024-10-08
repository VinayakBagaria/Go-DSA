package main

import (
	"go-dsa/binarysearch"
	"go-dsa/codechef"
	"go-dsa/dp"
	"go-dsa/graph"
	"go-dsa/heap"
	"go-dsa/leetcode"
	"go-dsa/linkedlist"
	"go-dsa/queue"
	"go-dsa/recursion"
	"go-dsa/sorting"
)

const decision = "binary_search"

func main() {
	switch decision {
	case "sorting":
		sorting.DoWork()
	case "linkedlist":
		linkedlist.DoWork()
	case "queue":
		queue.DoWork()
	case "graph":
		graph.DoWork()
	case "heap":
		heap.DoWork()
	case "leetcode":
		leetcode.DoWork()
	case "dynamic_programming":
		dp.DoWork()
	case "codechef":
		codechef.DoWork()
	case "recursion":
		recursion.DoWork()
	case "binary_search":
		binarysearch.DoWork()
	}
}
