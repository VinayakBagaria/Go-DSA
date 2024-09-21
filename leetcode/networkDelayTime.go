// https://leetcode.com/problems/network-delay-time/description/
package leetcode

import (
	"fmt"
	"math"
)

type Edge struct {
	v    int
	time int
}

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make(map[int][]Edge)
	for _, time := range times {
		u := time[0]
		v := time[1]
		adj[u] = append(adj[u], Edge{v, time[2]})
	}

	queue := []int{k}
	seen := make(map[int]int)
	for i := 1; i <= n; i++ {
		seen[i] = math.MaxInt
	}
	seen[k] = 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		t := seen[u]

		for _, e := range adj[u] {
			newTime := t + e.time
			if newTime < seen[e.v] {
				seen[e.v] = newTime
				queue = append(queue, e.v)
			}
		}
	}

	minTime := 0
	for _, v := range seen {
		minTime = max(minTime, v)
	}
	if minTime == math.MaxInt {
		return -1
	}
	return minTime
}

func DoNetworkDelayTime() {
	fmt.Println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}, {2, 3, 7}, {1, 3, 4}, {2, 1, 2}}, 3, 2))
}
