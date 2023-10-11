// https://leetcode.com/problems/number-of-flowers-in-full-bloom/description/
package random

import (
	"fmt"
	"sort"
)

func fullBloomFlowers(flowers [][]int, people []int) []int {
	starts := make([]int, len(flowers))
	ends := make([]int, len(flowers))

	for i, fl := range flowers {
		starts[i] = fl[0]
		ends[i] = fl[1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	result := make([]int, len(people))
	for i, person := range people {
		bloomed := upperBound(starts, person)
		died := lowerBound(ends, person)

		result[i] = bloomed - died
	}

	return result
}

func upperBound(arr []int, target int) int {
	l := 0
	r := len(arr)

	for l < r {
		mid := l + (r-l)/2
		if arr[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func lowerBound(arr []int, target int) int {
	l := 0
	r := len(arr)

	for l < r {
		mid := l + (r-l)/2
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func DoNumberOfFlowersInFullBloom() {
	fmt.Println(fullBloomFlowers([][]int{
		{1, 6},
		{3, 7},
		{9, 12},
		{4, 13},
	}, []int{2, 3, 7, 11}))
}
