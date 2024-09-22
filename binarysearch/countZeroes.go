// https://www.geeksforgeeks.org/problems/count-the-zeros2550/1
package binarysearch

import "fmt"

func countZeroes(arr []int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] == 1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return len(arr) - l
}

func DoCountZeroes() {
	fmt.Println(countZeroes([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}))
	fmt.Println(countZeroes([]int{0, 0, 0}))
}
