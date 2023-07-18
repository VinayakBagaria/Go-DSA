// https://leetcode.com/problems/smallest-sufficient-team/description/
package random

import (
	"fmt"
)

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	sCount := len(req_skills)
	pCount := len(people)

	skillIndex := make(map[string]int)
	for i, skill := range req_skills {
		skillIndex[skill] = i
	}

	peopleMask := make([]int, pCount)
	for i, personSkills := range people {
		for _, skill := range personSkills {
			peopleMask[i] |= (1 << skillIndex[skill])
		}
	}

	smallest := []int{}
	dp := make(map[string]int)

	var dfs func(int, []int, int)
	dfs = func(index int, ans []int, maskTillNow int) {
		if index == pCount {
			if maskTillNow == (1<<sCount)-1 {
				if len(smallest) == 0 || len(ans) < len(smallest) {
					smallest = append([]int{}, ans...)
				}
			}
			return
		}

		key := fmt.Sprintf("%d-%d", index, maskTillNow)
		if val, ok := dp[key]; ok {
			if val <= len(ans) {
				return
			}
		}

		// current index not coming
		dfs(index+1, ans, maskTillNow)
		// current index coming
		dfs(index+1, append(ans, index), maskTillNow|peopleMask[index])

		dp[key] = len(ans)
	}

	dfs(0, []int{}, 0)
	return smallest
}

func DoSmallestSufficientTeam() {
	fmt.Println(smallestSufficientTeam([]string{"java", "nodejs", "reactjs"}, [][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}}))
}
