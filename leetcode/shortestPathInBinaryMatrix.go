// https://leetcode.com/problems/shortest-path-in-binary-matrix/description/
package leetcode

import "fmt"

var ALL_DIRECTIONS = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}

	queue := []Cell{Cell{0, 0}}
	length := 0

	for len(queue) > 0 {
		size := len(queue)
		length++

		for i := 0; i < size; i++ {
			cell := queue[i]
			if cell.i == n-1 && cell.j == n-1 {
				return length
			}

			for _, d := range ALL_DIRECTIONS {
				newX := cell.i + d[0]
				newY := cell.j + d[1]

				if isOutside(newX, newY, n, n) || grid[newX][newY] != 0 {
					continue
				}
				grid[newX][newY] = 1
				queue = append(queue, Cell{newX, newY})
			}
		}

		queue = queue[size:]
	}

	return -1
}

func DoShortestPathInBinaryMatrix() {
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}))
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}))
}
