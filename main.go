package main

import (
	"go-dsa/dp"
	"go-dsa/graph"
	"go-dsa/heap"
	"go-dsa/linkedlist"
	"go-dsa/queue"
	"go-dsa/random"
	"go-dsa/sorting"
)

const decision = "graph"

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
	case "random":
		random.DoWork()
	case "dynamic_programming":
		dp.DoWork()
	}
}
