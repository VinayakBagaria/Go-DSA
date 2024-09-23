// https://www.geeksforgeeks.org/problems/rat-in-a-maze-problem/1

package graph

import "fmt"

func findPath(m [][]int) []string {
	result := []string{}
	visited := map[string]bool{}
	dirs := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	names := []string{"D", "U", "R", "L"}

	r := len(m)
	c := len(m[0])

	var solve func(int, int, string)
	solve = func(i, j int, current string) {
		if m[i][j] == 0 {
			return
		}

		if i == r-1 && j == c-1 {
			result = append(result, current)
			return
		}

		key := fmt.Sprintf("%d-%d", i, j)

		if _, ok := visited[key]; ok {
			return
		}

		visited[key] = true

		for k, d := range dirs {
			newI := i + d[0]
			newJ := j + d[1]

			if newI < 0 || newI >= r || newJ < 0 || newJ >= c {
				continue
			}

			solve(newI, newJ, current+names[k])
		}

		delete(visited, key)
	}

	solve(0, 0, "")
	return result
}

func DoRatInAMaze() {
	fmt.Println(findPath([][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	}))
}
