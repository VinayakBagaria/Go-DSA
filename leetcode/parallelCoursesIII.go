// https://leetcode.com/problems/parallel-courses-iii/
package leetcode

import "fmt"

func minimumTime(n int, relations [][]int, time []int) int {
	adjList := make(map[int][]int)
	indegree := make([]int, n+1)
	for _, rel := range relations {
		u := rel[0]
		v := rel[1]
		adjList[u] = append(adjList[u], v)
		indegree[v]++
	}

	queue := []int{}
	seen := make([]int, n+1)
	for u, indeg := range indegree {
		if u == 0 {
			continue
		}
		if indeg == 0 {
			queue = append(queue, u)
			seen[u] = time[u-1]
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		currTime := seen[u]

		for _, v := range adjList[u] {
			indegree[v]--
			newTime := currTime + time[v-1]
			seen[v] = max(newTime, seen[v])
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return max(seen...)
}

func DoParallelCoursesIII() {
	fmt.Println(minimumTime(3, [][]int{
		{1, 3},
		{2, 3},
	}, []int{2, 3, 5}))
	fmt.Println(minimumTime(5, [][]int{
		{1, 5},
		{2, 5},
		{3, 5},
		{3, 4},
		{4, 5},
	}, []int{1, 2, 3, 4, 5}))
}
