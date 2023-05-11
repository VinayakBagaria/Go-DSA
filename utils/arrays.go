package utils

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxOfArray(nums []int) int {
	maximum := nums[0]
	for _, num := range nums {
		maximum = Max(maximum, num)
	}
	return maximum
}

func CopyFromOneArrayToAnother(source []int) []int {
	var destination []int
	destination = append(destination, source...)
	return destination
}
