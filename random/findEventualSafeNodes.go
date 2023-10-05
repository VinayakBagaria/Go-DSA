// https://leetcode.com/problems/find-eventual-safe-states/description/
package random

import "fmt"

func eventualSafeNodesDfs(graph [][]int) []int {
	visited := make(map[int]struct{})
	inRecursion := make(map[int]struct{})

	var isCycle func(int) bool
	isCycle = func(u int) bool {
		visited[u] = struct{}{}
		inRecursion[u] = struct{}{}

		for _, v := range graph[u] {
			if exists(visited, v) {
				if exists(inRecursion, v) {
					return true
				}
			} else {
				if isCycle(v) {
					return true
				}
			}
		}

		delete(inRecursion, u)
		return false
	}

	for u := range graph {
		if !exists(visited, u) {
			isCycle(u)
		}
	}

	result := []int{}
	for u := range graph {
		if !exists(inRecursion, u) {
			result = append(result, u)
		}
	}

	return result
}

func exists(mapper map[int]struct{}, key int) bool {
	_, ok := mapper[key]
	return ok
}

func DoFindEventualSafeNodes() {
	fmt.Println(eventualSafeNodesDfs([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
	fmt.Println(eventualSafeNodesDfs([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}))
}
