// https://leetcode.com/problems/last-day-where-you-can-still-cross/description/
package random

import "fmt"

var directions = [][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type Point struct {
	row int
	col int
}

func isBounded(r, c, rows, cols int) bool {
	return r >= 0 && r < rows && c >= 0 && c < cols
}

func canCrossOnDay(row, col, day int, cells [][]int) bool {
	seen := make(map[Point]bool)

	// Till the day, fill those cells with water
	for i := 0; i < day; i++ {
		waterPoint := Point{row: cells[i][0] - 1, col: cells[i][1] - 1}
		seen[waterPoint] = true
	}

	// For 0th row, add all the land positions to to queue for BFS
	// queue element -> [row, col]
	queue := [][2]int{}
	for c := 0; c < col; c++ {
		landPoint := Point{row: 0, col: c}
		if _, ok := seen[landPoint]; !ok {
			queue = append(queue, [2]int{0, c})
			seen[landPoint] = true
		}
	}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			current := queue[i]
			r := current[0]
			c := current[1]
			if r == row-1 {
				return true
			}

			for _, d := range directions {
				newR := r + d[0]
				newC := c + d[1]

				if !isBounded(newR, newC, row, col) {
					continue
				}

				landPoint := Point{row: newR, col: newC}
				if _, ok := seen[landPoint]; !ok {
					seen[landPoint] = true
					queue = append(queue, [2]int{newR, newC})
				}
			}
		}

		queue = queue[size:]
	}

	return false
}

func latestDayToCross(row int, col int, cells [][]int) int {
	left := 1
	right := row * col
	ans := 0

	for left <= right {
		mid := left + (right-left)/2
		if canCrossOnDay(row, col, mid, cells) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func DoLastDayToCross() {
	// fmt.Println(latestDayToCross(2, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}))
	fmt.Println(latestDayToCross(3, 3, [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}}))
}
