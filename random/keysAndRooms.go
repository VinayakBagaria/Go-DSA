// https://leetcode.com/problems/keys-and-rooms/description/
package random

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	adj := make(map[int][]int)
	for u, visits := range rooms {
		for _, v := range visits {
			adj[u] = append(adj[u], v)
		}
	}

	seen := make(map[int]bool)
	queue := []int{0}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		seen[u] = true

		for _, v := range adj[u] {
			if _, ok := seen[v]; ok {
				continue
			}
			queue = append(queue, v)
		}
	}

	return len(seen) == len(rooms)
}

func DoKeysAndRooms() {
	fmt.Println(canVisitAllRooms([][]int{
		{1},
		{2},
		{3},
		{},
	}))
	fmt.Println(canVisitAllRooms([][]int{
		{1, 3},
		{3, 0, 1},
		{2},
		{0},
	}))
}
