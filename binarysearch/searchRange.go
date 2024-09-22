// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array
package binarysearch

import "fmt"

func firstTime(arr []int, x int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] >= x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func lastTime(arr []int, x int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] <= x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return r
}

func searchRange(nums []int, target int) []int {
	f := firstTime(nums, target)
	l := lastTime(nums, target)

	if f > l {
		return []int{-1, -1}
	}

	return []int{f, l}
}

func DoSearchRange() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{1, 1, 2, 2, 2, 2, 3}, 2))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}
