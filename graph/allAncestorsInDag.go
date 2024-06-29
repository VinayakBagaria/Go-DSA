// https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/
package graph

import (
	"fmt"
	"sort"
)

func getAllAncestorsInDag(n int, edges [][]int) [][]int {
	adjList := make(map[int][]int)
	indegree := make([]int, n+1)

	for _, e := range edges {
		u := e[0]
		v := e[1]
		adjList[u] = append(adjList[u], v)
		indegree[v]++
	}

	queue := []int{}
	for u, indeg := range indegree {
		if indeg == 0 {
			queue = append(queue, u)
		}
	}

	ancestors := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		ancestors[i] = make(map[int]bool)
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for _, v := range adjList[u] {
			indegree[v]--

			ancestors[v][u] = true
			for ancestor := range ancestors[u] {
				ancestors[v][ancestor] = true
			}

			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, 0, len(ancestors[i]))
		for ancestor := range ancestors[i] {
			ans[i] = append(ans[i], ancestor)
		}
		sort.Ints(ans[i])
	}

	return ans
}

func DoGetAllAncestorsInDag() {
	fmt.Println(getAllAncestorsInDag(8, [][]int{
		{0, 3},
		{0, 4},
		{1, 3},
		{2, 4},
		{2, 7},
		{3, 5},
		{3, 6},
		{3, 7},
		{4, 6},
	}))
}
