// https://leetcode.com/problems/course-schedule/description/
package random

import "fmt"

func canFinishDfs(numCourses int, prerequisites [][]int) bool {
	visited := make(map[int]struct{})
	inRecursion := make(map[int]struct{})
	adj := make(map[int][]int)

	for _, preq := range prerequisites {
		u := preq[1]
		v := preq[0]
		adj[u] = append(adj[u], v)
	}

	var isCycle func(int) bool
	isCycle = func(u int) bool {
		visited[u] = struct{}{}
		inRecursion[u] = struct{}{}

		for _, v := range adj[u] {
			if _, ok := visited[v]; ok {
				if _, ok = inRecursion[v]; ok {
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

	for u := 0; u < numCourses; u++ {
		if _, ok := visited[u]; !ok {
			if isCycle(u) {
				return false
			}
		}
	}

	return true
}

func canFinishBfs(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int)
	indegree := make(map[int]int)

	for _, preq := range prerequisites {
		u := preq[1]
		v := preq[0]
		adj[u] = append(adj[u], v)
		indegree[v]++
	}

	queue := []int{}

	for u := 0; u < numCourses; u++ {
		if indegree[u] == 0 {
			queue = append(queue, u)
		}
	}

	visitedCount := 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		visitedCount++

		for _, v := range adj[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return visitedCount == numCourses
}

func DoCourseSchedule() {
	fmt.Println("DFS:", canFinishDfs(2, [][]int{{1, 0}}))
	fmt.Println("BFS:", canFinishBfs(2, [][]int{{1, 0}}))
	fmt.Println("-------------")
	fmt.Println("DFS:", canFinishDfs(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println("BFS:", canFinishBfs(2, [][]int{{1, 0}, {0, 1}}))
}
