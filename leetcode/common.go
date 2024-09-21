package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(vars ...int) int {
	result := vars[0]
	for _, v := range vars {
		if result < v {
			result = v
		}
	}
	return result
}

func min(vars ...int) int {
	result := vars[0]
	for _, v := range vars {
		if v < result {
			result = v
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func reverse(str string) string {
	result := ""
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

const MOD = 1e9 + 7
