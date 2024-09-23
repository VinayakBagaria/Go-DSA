package recursion

import "fmt"

func merge(a1, a2 []int) []int {
	merged := []int{}
	i := 0
	j := 0

	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			merged = append(merged, a1[i])
			i++
		} else {
			merged = append(merged, a2[j])
			j++
		}
	}

	for i < len(a1) {
		merged = append(merged, a1[i])
		i++
	}

	for j < len(a2) {
		merged = append(merged, a2[j])
		j++
	}

	return merged
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	a1 := mergeSort(arr[:mid])
	a2 := mergeSort(arr[mid:])
	return merge(a1, a2)
}

func DoMergeSort() {
	fmt.Println(mergeSort([]int{3, 1, 2}))
}
