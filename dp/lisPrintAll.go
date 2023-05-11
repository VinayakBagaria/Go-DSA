// https://www.youtube.com/watch?v=3mD20VSib5E
package dp

import (
	"fmt"
	"go-dsa/utils"
)

func doLis(nums []int) ([]int, int) {
	dp := make([]int, len(nums))
	dp[0] = 1
	overallMax := 0

	for i, num := range nums {
		maximum := 0
		for j := 0; j < i; j++ {
			if nums[j] < num {
				maximum = utils.Max(maximum, dp[j])
			}
		}

		dp[i] = maximum + 1
		if dp[i] > overallMax {
			overallMax = dp[i]
		}
	}

	return dp, overallMax
}

type Point struct {
	index      int
	lisAtIndex int
	pathSoFar  []int
}

func getAllCombinations(nums []int) [][]int {
	dp, overallMax := doLis(nums)

	var queue []Point
	for i, dpValue := range dp {
		if dpValue == overallMax {
			queue = append(queue, Point{
				index:      i,
				lisAtIndex: dp[i],
				pathSoFar:  []int{nums[i]},
			})
		}
	}

	var paths [][]int

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]

		if popped.lisAtIndex == 1 {
			paths = append(paths, popped.pathSoFar)
			continue
		}

		currentIndex := popped.index
		for i := currentIndex; i >= 0; i-- {
			value := nums[i]
			pathSoFar := append([]int{value}, popped.pathSoFar...)

			if dp[i]+1 == dp[currentIndex] && value <= nums[popped.index] {

				queue = append(queue, Point{
					index:      i,
					lisAtIndex: dp[i],
					pathSoFar:  pathSoFar,
				})
			}
		}
	}

	return paths
}

func DoLisPrintAll() {
	fmt.Println("Print all LIS")
	inputs := [][]int{
		// {1, 2, 4, 3},
		// {10, 22, 9, 33, 21, 50, 41, 60, 80, 3},
		{10, 22, 42, 33, 21, 50, 41, 60, 80, 3},
	}

	for _, input := range inputs {
		fmt.Printf("Input: %v\nOutput: %v\n", input, getAllCombinations(input))
	}
}
