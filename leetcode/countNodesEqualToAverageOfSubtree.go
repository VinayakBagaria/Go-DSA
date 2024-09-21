// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/description/
package leetcode

import "fmt"

func averageOfSubtree(root *TreeNode) int {
	count := 0

	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		sum := node.Val
		lSum, lCount := dfs(node.Left)
		rSum, rCount := dfs(node.Right)

		totalSum := sum + lSum + rSum
		totalCount := 1 + lCount + rCount

		if totalSum/totalCount == node.Val {
			count++
		}

		return totalSum, totalCount
	}

	dfs(root)
	return count
}

func DoCountNodesEqualToAverageOfSubtree() {
	fmt.Println(averageOfSubtree(&TreeNode{1, nil, nil}))
}
