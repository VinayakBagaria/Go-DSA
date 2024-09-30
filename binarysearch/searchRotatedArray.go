// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
package binarysearch

import "fmt"

func search(arr []int, target int) int {
	s := 0
	e := len(arr) - 1

	for s <= e {
		mid := s + (e-s)/2
		if arr[mid] == target {
			return mid
		}

		// Left half of array sorted condition
		if arr[s] <= arr[mid] {
			if arr[s] <= target && target <= arr[mid] {
				e = mid - 1
			} else {
				s = mid + 1
			}
		} else {
			// Right half of array sorted condition
			if arr[mid] <= target && target <= arr[e] {
				s = mid + 1
			} else {
				e = mid - 1
			}
		}
	}

	return -1
}

func DoSearchRotatedArray() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
