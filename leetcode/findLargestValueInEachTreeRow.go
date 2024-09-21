// https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/
package leetcode

import (
	"fmt"
	"math"
)

func largestValuesDfs(root *TreeNode) []int {
	result := []int{}

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level == len(result) {
			result = append(result, node.Val)
		} else {
			result[level] = max(result[level], node.Val)
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)
	return result
}

func largestValuesBfs(root *TreeNode) []int {
	result := []int{}
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		size := len(queue)
		maxVal := math.MinInt

		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Val > maxVal {
				maxVal = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]
		result = append(result, maxVal)
	}

	return result
}

func DoFindLargestValueInEachTreeRow() {
	fmt.Println(largestValuesDfs(&TreeNode{1, nil, nil}))
	fmt.Println(largestValuesBfs(&TreeNode{1, nil, nil}))
}
