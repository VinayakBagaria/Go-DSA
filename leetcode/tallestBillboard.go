// https://leetcode.com/problems/tallest-billboard/description/
package leetcode

import "fmt"

func copyMap[K, V comparable](old map[K]V) map[K]V {
	newMap := make(map[K]V)
	for k, v := range old {
		newMap[k] = v
	}
	return newMap
}

func tallestBillboard(rods []int) int {
	dp := make(map[int]int)
	dp[0] = 0

	for _, r := range rods {
		newDp := copyMap(dp)

		for diff, taller := range dp {
			shorter := taller - diff

			// Add to taller
			newDp[diff+r] = max(newDp[diff+r], taller+r)

			// Add to shorter
			newDiff := abs(taller - (shorter + r))
			newDp[newDiff] = max(newDp[newDiff], shorter+r, taller)
		}

		dp = newDp
	}

	return dp[0]
}

func DoTallesBillboard() {
	fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
}
