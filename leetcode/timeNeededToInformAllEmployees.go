// https://leetcode.com/problems/time-needed-to-inform-all-employees/description/
package leetcode

import "fmt"

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	adj := make(map[int][]int)
	for i, man := range manager {
		adj[man] = append(adj[man], i)
	}

	queue := []int{headID}
	taken := make([]int, n+1)
	taken[headID] = informTime[headID]
	maxVal := informTime[headID]

	for len(queue) > 0 {
		person := queue[0]
		queue = queue[1:]
		currTaken := taken[person]

		for _, sub := range adj[person] {
			queue = append(queue, sub)
			taken[sub] = currTaken + informTime[sub]
			if maxVal < taken[sub] {
				maxVal = taken[sub]
			}
		}
	}

	return maxVal
}

func DoTimeNeededToInformAllEmployees() {
	fmt.Println(numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}))
	fmt.Println(numOfMinutes(5, 0, []int{-1, 0, 0, 1, 2}, []int{0, 2, 3, 4, 7}))
}
