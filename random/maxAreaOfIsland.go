// https://leetcode.com/problems/max-area-of-island/description/
package random

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0
		result := 0
		for _, dir := range DIRECTIONS {
			newX := i + dir[0]
			newY := j + dir[1]
			result += dfs(newX, newY)
		}
		return 1 + result
	}

	maxVal := 0
	for i, row := range grid {
		for j := range row {
			maxVal = max(maxVal, dfs(i, j))
		}
	}

	return maxVal
}

func DoMaxAreaOfIsland() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}
