// https://leetcode.com/problems/find-the-town-judge/description/
package leetcode

import "fmt"

func findJudge(n int, trust [][]int) int {
	outdegree := make([]int, n+1)
	indegree := make([]int, n+1)
	for _, group := range trust {
		u := group[0]
		v := group[1]
		outdegree[u]++
		indegree[v]++
	}

	for u := 1; u <= n; u++ {
		// outdegree 0 -> trusts no-one
		// indegree n - 1 -> trusted by everyone except himself
		if outdegree[u] == 0 && indegree[u] == n-1 {
			return u
		}
	}

	return -1
}

func DoFindTheTownJudge() {
	fmt.Println(findJudge(2, [][]int{
		{1, 2},
	}))
	fmt.Println(findJudge(3, [][]int{
		{1, 3},
		{2, 3},
	}))
	fmt.Println(findJudge(3, [][]int{
		{1, 3},
		{2, 3},
		{3, 1},
	}))
}
