// https://leetcode.com/problems/redundant-connection/description/
package leetcode

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
		rank[i] = 1
	}

	var find func(u int) int
	find = func(u int) int {
		if parent[u] == u {
			return u
		}
		parent[u] = find(parent[u])
		return parent[u]
	}

	var union = func(u, v int) bool {
		p1 := find(u)
		p2 := find(v)

		if p1 == p2 {
			return false
		}

		if rank[p1] > rank[p2] {
			parent[p2] = p1
		} else if rank[p1] < rank[p2] {
			parent[p1] = p2
		} else {
			rank[p1]++
			parent[p2] = p1
		}

		return true
	}

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		if !union(u, v) {
			return edge
		}
	}

	return []int{}
}

func DoRedundantConnection() {
	fmt.Println(findRedundantConnection([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 4},
		{1, 5},
	}))
}
