// https://leetcode.com/problems/majority-element-ii/
// Boyer-Moore algorithm
package leetcode

import "fmt"

func majorityElement2(nums []int) []int {
	desired := (len(nums) / 3) + 1
	candidate1 := 0
	count1 := 0

	candidate2 := 0
	count2 := 0

	for _, n := range nums {
		if candidate1 == n {
			count1++
		} else if candidate2 == n {
			count2++
		} else if count1 == 0 {
			candidate1 = n
			count1 = 1
		} else if count2 == 0 {
			candidate2 = n
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	freq1 := 0
	freq2 := 0

	for _, n := range nums {
		if n == candidate1 {
			freq1++
		} else if n == candidate2 {
			freq2++
		}
	}

	result := []int{}
	if freq1 >= desired {
		result = append(result, candidate1)
	}
	if freq2 >= desired {
		result = append(result, candidate2)
	}

	return result
}

func DoMajorityElement2() {
	fmt.Println(majorityElement2([]int{3, 2, 3}))
	fmt.Println(majorityElement2([]int{1}))
	fmt.Println(majorityElement2([]int{1, 2}))
}
