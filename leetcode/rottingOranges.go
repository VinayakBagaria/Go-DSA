// https://leetcode.com/problems/rotting-oranges/description/
package leetcode

import "fmt"

var ROTTEN = 2
var FRESH = 1

type Cell struct {
	i, j int
}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	freshCount := 0

	queue := []Cell{}
	for i, row := range grid {
		for j, orange := range row {
			if orange == ROTTEN {
				queue = append(queue, Cell{i, j})
			} else if orange == FRESH {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	time := 0
	for len(queue) > 0 {
		size := len(queue)
		time++
		for i := 0; i < size; i++ {
			cell := queue[i]
			for _, d := range DIRECTIONS {
				newX := cell.i + d[0]
				newY := cell.j + d[1]
				if isOutside(newX, newY, m, n) || grid[newX][newY] != FRESH {
					continue
				}
				grid[newX][newY] = ROTTEN
				queue = append(queue, Cell{newX, newY})
				freshCount--
			}
		}
		queue = queue[size:]
	}

	if freshCount == 0 {
		return time - 1
	}
	return -1
}

func isOutside(i, j, m, n int) bool {
	return i < 0 || i >= m || j < 0 || j >= n
}

func DoRottingOranges() {
	fmt.Println(orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}))
	fmt.Println(orangesRotting([][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}))
	fmt.Println(orangesRotting([][]int{
		{0, 2},
	}))
}
