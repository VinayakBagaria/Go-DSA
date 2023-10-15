package graph

import "fmt"

func isCycle(connections [][]int, n int) bool {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}

	var find func(int) int
	find = func(u int) int {
		if parent[u] == u {
			return u
		}
		parent[u] = find(parent[u])
		return parent[u]
	}

	var union = func(u, v int) {
		p1 := find(u)
		p2 := find(v)

		if rank[p1] > rank[p2] {
			parent[p2] = p1
		} else if rank[p1] < rank[p2] {
			parent[p1] = p2
		} else {
			rank[p1]++
			parent[p2] = p1
		}
	}

	adj := make(map[int][]int)
	for _, conn := range connections {
		u := conn[0]
		v := conn[1]
		adj[u] = append(adj[u], v)
	}

	for u, neighbours := range adj {
		for _, v := range neighbours {
			union(u, v)
			union(v, u)
		}
	}

	first := parent[0]
	for _, p := range parent {
		if p != first {
			return false
		}
	}

	return true
}

func DoUnion() {
	connections := [][]int{
		{2, 0},
		{3, 0},
		{4, 1},
		{4, 2},
		{4, 0},
	}
	fmt.Println(isCycle(connections, 5))
}
