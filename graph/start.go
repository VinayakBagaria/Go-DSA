package graph

import "fmt"

const separator = "--------------------------------------"

func DoWork() {
	fmt.Println("Edge array => Adjancency List")
	fmt.Println(MakeAdjacencyList(edges))
	fmt.Println(separator)
	DoTraversal()
	fmt.Println(separator)
	DoHasPathDirected()
	fmt.Println(separator)
	DoHasPathUndirected()
	fmt.Println(separator)
	DoConnectedComponents()
	fmt.Println(separator)
	DoLargestComponentCount()
	fmt.Println(separator)
	DoShortestPathDistance()
	fmt.Println(separator)
	DoDistinctIslandCount()
	fmt.Println(separator)
	DoHasCycleDirected()
	fmt.Println(separator)
	DoHasCycleUndirected()
	fmt.Println(separator)
	DoTopologicalSort()
	fmt.Println(separator)
	DoWordLadder()
}
