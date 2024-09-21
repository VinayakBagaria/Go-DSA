package recursion

import "fmt"

func reverseArr(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	return append(reverseArr(arr[1:]), arr[0])
}

func DoReverseArray() {
	arr := []int{4, 5, 7}
	fmt.Println(reverseArr(arr))
}
