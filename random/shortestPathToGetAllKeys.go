package random

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isLock(element byte) bool {
	return element >= 'A' && element <= 'Z'
}

func isKey(element byte) bool {
	return element >= 'a' && element <= 'z'
}

func hasMoreKeys(multiGrid [][]byte) bool {
	for _, row := range multiGrid {
		for _, element := range row {
			if isKey(element) {
				return true
			}
		}
	}
	return false
}

func shortestPathAllKeys(grid []string) int {
	multiGrid := [][]byte{}
	src := [2]int{-1, -1}
	totalKeys := 0
	for i, row := range grid {
		values := []byte{}
		for j := range row {
			values = append(values, row[j])
			if row[j] == '@' {
				src = [2]int{i, j}
			}
			if isKey(row[j]) {
				totalKeys++
			}
		}
		multiGrid = append(multiGrid, values)
	}

	minValue := math.MaxInt32

	var dfs func(int, int, int, int, int)
	dfs = func(r, c, steps, keyCount, encountered int) {
		if r < 0 || r >= len(multiGrid) || c < 0 || c >= len(multiGrid[0]) {
			return
		}

		val := multiGrid[r][c]

		if val == '#' {
			return
		}

		if isKey(val) {
			keyCount++
			encountered++
		}
		if isLock(val) {
			if keyCount == 0 {
				return
			}
			keyCount--
		}

		multiGrid[r][c] = '#'

		if encountered == totalKeys || !hasMoreKeys(multiGrid) {
			minValue = min(minValue, steps)
			multiGrid[r][c] = val
			return
		}

		dfs(r+1, c, steps+1, keyCount, encountered)
		dfs(r-1, c, steps+1, keyCount, encountered)
		dfs(r, c+1, steps+1, keyCount, encountered)
		dfs(r, c-1, steps+1, keyCount, encountered)

		multiGrid[r][c] = val
	}

	dfs(src[0], src[1], 0, 0, 0)
	if minValue == math.MaxInt32 {
		return -1
	}
	return minValue
}

func DoShortestPathToGetAllKeys() {
	fmt.Println(shortestPathAllKeys([]string{"@.a..", "###.#", "b.A.B"}))
}
