// https://leetcode.com/problems/find-eventual-safe-states/description/
package leetcode

import (
	"fmt"
)

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	safe := make(map[int]bool)
	seen := make(map[int]bool)

	var dfs func(int) bool
	dfs = func(u int) bool {
		if val, ok := safe[u]; ok {
			return val
		}

		if _, ok := seen[u]; ok {
			return false
		}

		seen[u] = true
		for _, v := range graph[u] {
			if !dfs(v) {
				safe[u] = false
				return safe[u]
			}
		}

		safe[u] = true
		return safe[u]
	}

	nodes := []int{}
	for i := 0; i < n; i++ {
		if dfs(i) {
			nodes = append(nodes, i)
		}
	}
	return nodes
}

func DoFindEventualSafeNodes() {
	fmt.Println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
	fmt.Println(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}))
}
