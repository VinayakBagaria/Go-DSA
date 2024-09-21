package codechef

import "fmt"

func mainFunction(arr []int) []int {
	reduced := []int{}
	reduced = append(reduced, arr[0])

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			reduced = append(reduced, arr[i])
		}
	}

	return reduced
}

func DoRemoveDuplicates() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCases > 0 {
		testCases--

		var n int
		fmt.Scanln(&n)
		numbers := ScanArray()

		result := mainFunction(numbers)
		fmt.Println(len(result))
		fmt.Println(PrintArray(result))
	}
}
